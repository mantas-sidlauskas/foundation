// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubeops

import (
	"context"
	"encoding/json"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	k8s "k8s.io/client-go/kubernetes"
	"namespacelabs.dev/foundation/engine/ops"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/runtime/kubernetes/client"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubedef"
	kobs "namespacelabs.dev/foundation/runtime/kubernetes/kubeobserver"
	"namespacelabs.dev/foundation/runtime/kubernetes/networking/ingress"
	fnschema "namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/schema/orchestration"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func registerApply() {
	ops.RegisterVFuncs(ops.VFuncs[*kubedef.OpApply, *parsedApply]{
		Parse: func(ctx context.Context, def *fnschema.SerializedInvocation, apply *kubedef.OpApply) (*parsedApply, error) {
			if apply.BodyJson == "" {
				return nil, fnerrors.InternalError("apply.Body is required")
			}

			var parsed unstructured.Unstructured
			if err := json.Unmarshal([]byte(apply.BodyJson), &parsed); err != nil {
				return nil, fnerrors.BadInputError("kubernetes.apply: failed to parse resource: %w", err)
			}

			return &parsedApply{obj: &parsed, spec: apply}, nil
		},

		Handle: func(ctx context.Context, d *fnschema.SerializedInvocation, spec *parsedApply) (*ops.HandleResult, error) {
			return apply(ctx, d.Description, fnschema.PackageNames(d.Scope...), spec)
		},

		PlanOrder: func(apply *parsedApply) (*fnschema.ScheduleOrder, error) {
			return kubedef.PlanOrder(apply.obj.GroupVersionKind()), nil
		},
	})
}

type parsedApply struct {
	obj  *unstructured.Unstructured
	spec *kubedef.OpApply
}

func apply(ctx context.Context, desc string, scope []fnschema.PackageName, apply *parsedApply) (*ops.HandleResult, error) {
	gv := apply.obj.GroupVersionKind().GroupVersion()
	if gv.Version == "" {
		return nil, fnerrors.InternalError("%s: APIVersion is required", desc)
	}

	cluster, err := kubedef.InjectedKubeCluster(ctx)
	if err != nil {
		return nil, err
	}

	restcfg := cluster.PreparedClient().RESTConfig

	var resource *schema.GroupVersionResource
	var res unstructured.Unstructured

	action := tasks.Action("kubernetes.apply").
		Scope(scope...).
		HumanReadablef(desc).
		Arg("name", apply.obj.GetName())

	if ns := apply.obj.GetNamespace(); ns != "" {
		action = action.Arg("namespace", ns)
	}

	if err := action.RunWithOpts(ctx, tasks.RunOpts{
		Wait: func(ctx context.Context) (bool, error) {
			var err error
			resource, err = resolveResource(ctx, cluster, apply.obj.GroupVersionKind())
			if err != nil {
				return false, err
			}

			tasks.Attachments(ctx).AddResult("resource", resource.Resource)

			// Creating Deployments and Statefulsets that refer to the
			// default service account, requires that the default service
			// account actually exists. And creating the default service
			// account takes a bit of time after creating a namespace.
			waitOnNamespace, err := requiresWaitForNamespace(apply)
			if err != nil {
				return false, fnerrors.InternalError("failed to determine object namespace: %w", err)
			}

			if waitOnNamespace {
				if err := waitForDefaultServiceAccount(ctx, cluster.PreparedClient().Clientset, apply.obj.GetNamespace()); err != nil {
					return false, err
				}
			}

			if !cluster.PreparedClient().Configuration.Ephemeral {
				if apply.obj.GetAPIVersion() == "networking.k8s.io/v1" && apply.obj.GetKind() == "Ingress" {
					if err := ingress.EnsureState(ctx, cluster); err != nil {
						return false, err
					}
				}
			}

			return false, nil
		},
		Run: func(ctx context.Context) error {
			client, err := client.MakeGroupVersionBasedClient(ctx, restcfg, resource.GroupVersion())
			if err != nil {
				return err
			}

			patchOpts := kubedef.Ego().ToPatchOptions()
			req := client.Patch(types.ApplyPatchType)
			if apply.obj.GetNamespace() != "" {
				req = req.Namespace(apply.obj.GetNamespace())
			}

			prepReq := req.Resource(resource.Resource).
				Name(apply.obj.GetName()).
				VersionedParams(&patchOpts, metav1.ParameterCodec).
				Body([]byte(apply.spec.BodyJson))

			if OutputKubeApiURLs {
				fmt.Fprintf(console.Debug(ctx), "kubernetes: api patch call %q\n", prepReq.URL())
			}

			return prepReq.Do(ctx).Into(&res)
		}}); err != nil {
		return nil, fnerrors.InvocationError("%s: failed to apply: %w", desc, err)
	}

	if err := checkResetCRDCache(ctx, cluster, apply.obj.GroupVersionKind()); err != nil {
		return nil, err
	}

	if apply.obj.GetNamespace() == kubedef.AdminNamespace && !kubedef.HasFocusMark(apply.obj.GetLabels()) {
		// don't wait for changes to admin namespace, unless they are in focus
		return &ops.HandleResult{}, nil
	}

	if apply.spec.CheckGenerationCondition.GetType() != "" {
		generation, found1, err1 := unstructured.NestedInt64(res.Object, "metadata", "generation")
		if err1 != nil {
			return nil, fnerrors.InternalError("failed to wait on resource: %w", err)
		}
		if !found1 {
			return nil, fnerrors.InternalError("failed to wait on resource: no metadata.generation")
		}

		return &ops.HandleResult{Waiters: []ops.Waiter{kobs.WaitOnGenerationCondition{
			RestConfig:         restcfg,
			Namespace:          apply.obj.GetNamespace(),
			Name:               apply.obj.GetName(),
			ExpectedGeneration: generation,
			ConditionType:      apply.spec.CheckGenerationCondition.Type,
			Resource:           *resource,
		}.WaitUntilReady}}, nil
	}

	switch {
	case kubedef.IsDeployment(apply.obj), kubedef.IsStatefulSet(apply.obj):
		generation, found1, err1 := unstructured.NestedInt64(res.Object, "metadata", "generation")
		observedGen, found2, err2 := unstructured.NestedInt64(res.Object, "status", "observedGeneration")
		if err2 != nil || !found2 {
			observedGen = 0 // Assume no generation exists.
		}

		// XXX print a warning if expected fields are missing.
		if err1 == nil && found1 {
			var waiters []ops.Waiter
			for _, sc := range scope {
				w := kobs.WaitOnResource{
					RestConfig:       restcfg,
					Description:      desc,
					Namespace:        apply.obj.GetNamespace(),
					Name:             apply.obj.GetName(),
					GroupVersionKind: apply.obj.GroupVersionKind(),
					Scope:            sc,
					PreviousGen:      observedGen,
					ExpectedGen:      generation,
				}
				waiters = append(waiters, w.WaitUntilReady)
			}
			return &ops.HandleResult{
				Waiters: waiters,
			}, nil
		} else {
			fmt.Fprintf(console.Warnings(ctx), "missing generation data from %s: %v / %v [found1=%v found2=%v]\n",
				apply.obj.GetKind(), err1, err2, found1, found2)
		}

	case kubedef.IsPod(apply.obj):
		waiters := []ops.Waiter{func(ctx context.Context, ch chan *orchestration.Event) error {
			if ch != nil {
				defer close(ch)
			}

			return kobs.WaitForCondition(ctx, cluster.PreparedClient().Clientset, tasks.Action("pod.wait").Scope(scope...),
				kobs.WaitForPodConditition(
					apply.obj.GetNamespace(),
					kobs.PickPod(apply.obj.GetName()),
					func(ps v1.PodStatus) (bool, error) {
						meta, err := json.Marshal(ps)
						if err != nil {
							return false, fnerrors.InternalError("failed to marshal pod status: %w", err)
						}

						ev := &orchestration.Event{
							ResourceId:          fmt.Sprintf("%s/%s", apply.obj.GetNamespace(), apply.obj.GetName()),
							Kind:                apply.obj.GetKind(),
							Category:            "Servers deployed",
							Ready:               orchestration.Event_NOT_READY,
							ImplMetadata:        meta,
							RuntimeSpecificHelp: fmt.Sprintf("kubectl -n %s describe pod %s", apply.obj.GetNamespace(), apply.obj.GetName()),
						}

						// XXX this under-reports scope.
						if len(scope) > 0 {
							ev.Scope = scope[0].String()
						}

						ev.WaitStatus = append(ev.WaitStatus, kobs.WaiterFromPodStatus(apply.obj.GetNamespace(), apply.obj.GetName(), ps))

						var done bool
						if ps.Phase == v1.PodFailed || ps.Phase == v1.PodSucceeded {
							// If the pod is finished, then don't wait further.
							done = true
						} else {
							done, _ = kobs.MatchPodCondition(v1.PodReady)(ps)
							if done {
								ev.Ready = orchestration.Event_READY
							}
						}

						if ch != nil {
							ch <- ev
						}
						return done, nil
					}))
		}}

		return &ops.HandleResult{
			Waiters: waiters,
		}, nil
	}

	return nil, nil
}

func requiresWaitForNamespace(apply *parsedApply) (bool, error) {
	switch {
	case kubedef.IsDeployment(apply.obj):
		// XXX change to unstructured lookups.
		var d appsv1.Deployment
		if err := json.Unmarshal([]byte(apply.spec.BodyJson), &d); err != nil {
			return false, err
		}
		if d.Spec.Template.Spec.ServiceAccountName == "" || d.Spec.Template.Spec.ServiceAccountName == "default" {
			return true, nil
		}

	case kubedef.IsStatefulSet(apply.obj):
		var d appsv1.StatefulSet
		if err := json.Unmarshal([]byte(apply.spec.BodyJson), &d); err != nil {
			return false, err
		}
		if d.Spec.Template.Spec.ServiceAccountName == "" || d.Spec.Template.Spec.ServiceAccountName == "default" {
			return true, nil
		}

	case kubedef.IsPod(apply.obj):
		var d v1.Pod
		if err := json.Unmarshal([]byte(apply.spec.BodyJson), &d); err != nil {
			return false, err
		}
		if d.Spec.ServiceAccountName == "" || d.Spec.ServiceAccountName == "default" {
			return true, nil
		}
	}

	return false, nil
}

func waitForDefaultServiceAccount(ctx context.Context, c *k8s.Clientset, namespace string) error {
	return tasks.Action("kubernetes.apply.wait-for-namespace").Arg("name", namespace).Run(ctx, func(ctx context.Context) error {
		w, err := c.CoreV1().ServiceAccounts(namespace).Watch(ctx, metav1.ListOptions{})
		if err != nil {
			return fnerrors.InternalError("kubernetes: failed to wait until the namespace was ready: %w", err)
		}

		defer w.Stop()

		// Wait until the default service account has been created.
		for ev := range w.ResultChan() {
			if account, ok := ev.Object.(*v1.ServiceAccount); ok && ev.Type == watch.Added {
				if account.Name == "default" {
					return nil // Service account is ready.
				}
			}
		}

		return nil
	})
}
