// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package deploy

import (
	"context"
	"fmt"

	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/executor"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/planning"
	"namespacelabs.dev/foundation/internal/runtime"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/cfg"
	"namespacelabs.dev/foundation/std/tasks"
)

type ComputeIngressResult struct {
	Fragments []*schema.IngressFragment

	stack *schema.Stack
}

func DeferredIngress(planner planning.Planner, stack *schema.Stack) compute.Computable[[]*schema.IngressFragment] {
	return compute.Inline(tasks.Action("ingress.compute"), func(ctx context.Context) ([]*schema.IngressFragment, error) {
		return computeDeferredIngresses(ctx, planner.Context, planner.Runtime, stack)
	})
}

func MergeIngresses(block ...compute.Computable[[]*schema.IngressFragment]) compute.Computable[[]*schema.IngressFragment] {
	return compute.Merge("flatten", compute.Collect(tasks.Action("ingress.merge"), block...))
}

func MaybeAllocateDomainCertificates(ingresses compute.Computable[[]*schema.IngressFragment], env *schema.Environment, stack *schema.Stack, allocate bool) compute.Computable[*ComputeIngressResult] {
	return compute.Transform("allocate-domains", ingresses, func(ctx context.Context, allFragments []*schema.IngressFragment) (*ComputeIngressResult, error) {
		eg := executor.New(ctx, "compute.ingress")
		for _, fragment := range allFragments {
			fragment := fragment // Close fragment.

			eg.Go(func(ctx context.Context) error {
				sch := stack.GetServer(schema.PackageName(fragment.Owner))
				if sch == nil {
					return fnerrors.BadInputError("%s: not present in the stack", fragment.Owner)
				}

				if allocate {
					var err error
					fragment.DomainCertificate, err = runtime.MaybeAllocateDomainCertificate(ctx, env, sch, fragment.Domain)
					if err != nil {
						return err
					}
				}

				return nil
			})
		}

		if err := eg.Wait(); err != nil {
			return nil, err
		}

		return &ComputeIngressResult{
			stack:     stack,
			Fragments: allFragments,
		}, nil
	})
}

func PlanIngressDeployment(rc runtime.Planner, c compute.Computable[*ComputeIngressResult]) compute.Computable[*runtime.DeploymentPlan] {
	return compute.Transform("plan ingress", c, func(ctx context.Context, res *ComputeIngressResult) (*runtime.DeploymentPlan, error) {
		return rc.PlanIngress(ctx, res.stack, res.Fragments)
	})
}

func computeDeferredIngresses(ctx context.Context, env cfg.Context, planner runtime.Planner, stack *schema.Stack) ([]*schema.IngressFragment, error) {
	var fragments []*schema.IngressFragment

	// XXX parallelism.
	for _, srv := range stack.Entry {
		frags, err := runtime.ComputeIngress(ctx, env, planner, srv, stack.Endpoint)
		if err != nil {
			return nil, err
		}
		fragments = append(fragments, frags...)
	}

	return fragments, nil
}

func ingressesFromHandlerResult(def compute.Computable[*handlerResult]) compute.Computable[[]*schema.IngressFragment] {
	return compute.Transform("parse computed ingress", def, func(ctx context.Context, h *handlerResult) ([]*schema.IngressFragment, error) {
		var fragments []*schema.IngressFragment

		for _, computed := range h.MergedComputedConfigurations().GetEntry() {
			for _, conf := range computed.Configuration {
				p := &schema.IngressFragment{}
				if !conf.Impl.MessageIs(p) {
					continue
				}

				if err := conf.Impl.UnmarshalTo(p); err != nil {
					return nil, err
				}

				fmt.Fprintf(console.Debug(ctx), "%s: received domain: %+v\n", conf.Owner, p.Domain)

				fragments = append(fragments, p)
			}
		}

		return fragments, nil
	})
}

func computeIngressWithHandlerResult(planner planning.Planner, stack *planning.Stack, additional ...compute.Computable[[]*schema.IngressFragment]) compute.Computable[*ComputeIngressResult] {
	var all []compute.Computable[[]*schema.IngressFragment]
	all = append(all, DeferredIngress(planner, stack.Proto()))
	all = append(all, additional...)

	merged := MergeIngresses(all...)

	return MaybeAllocateDomainCertificates(merged, planner.Context.Environment(), stack.Proto(), AlsoDeployIngress)
}
