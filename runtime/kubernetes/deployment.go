// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package kubernetes

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	applycorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	applymetav1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/runtime/kubernetes/kubedef"
	"namespacelabs.dev/foundation/schema"
	"sigs.k8s.io/yaml"
)

const kubeNode schema.PackageName = "namespacelabs.dev/foundation/std/runtime/kubernetes"

type perEnvConf struct {
	dashnessPeriod        int32
	livenessInitialDelay  int32
	readinessInitialDelay int32
	probeTimeout          int32
	failureThreshold      int32
}

var constants = map[schema.Environment_Purpose]perEnvConf{
	schema.Environment_DEVELOPMENT: {
		dashnessPeriod:        1,
		livenessInitialDelay:  1,
		readinessInitialDelay: 1,
		probeTimeout:          1,
		failureThreshold:      3,
	},
	schema.Environment_TESTING: {
		dashnessPeriod:        1,
		livenessInitialDelay:  1,
		readinessInitialDelay: 1,
		probeTimeout:          1,
		failureThreshold:      3,
	},
	schema.Environment_PRODUCTION: {
		dashnessPeriod:        3,
		livenessInitialDelay:  1,
		readinessInitialDelay: 3,
		probeTimeout:          1,
		failureThreshold:      5,
	},
}

func getEnv(c *applycorev1.ContainerApplyConfiguration, name string) (string, bool) {
	for _, env := range c.Env {
		if *env.Name == name {
			return *env.Value, true
		}
	}

	return "", false
}

func getArg(c *applycorev1.ContainerApplyConfiguration, name string) (string, bool) {
	for _, arg := range c.Args {
		if !strings.HasPrefix(arg, "-") {
			continue
		}
		// Remove up to two dashes.
		cleaned := strings.TrimPrefix(strings.TrimPrefix(arg, "-"), "-")
		parts := strings.SplitN(cleaned, "=", 2)
		if len(parts) != 2 {
			continue
		}

		if parts[0] == name {
			return parts[1], true
		}
	}

	return "", false
}

func toProbe(endpoint *schema.InternalEndpoint, md *schema.ServiceMetadata) (*applycorev1.ProbeApplyConfiguration, error) {
	exported := &schema.HttpExportedService{}
	if err := md.Details.UnmarshalTo(exported); err != nil {
		return nil, fnerrors.InternalError("expected a HttpExportedService: %w", err)
	}

	return applycorev1.Probe().WithHTTPGet(
		applycorev1.HTTPGetAction().WithPath(exported.GetPath()).
			WithPort(intstr.FromInt(int(endpoint.GetPort().GetContainerPort())))), nil
}

type deployOpts struct {
}

func (r boundEnv) prepareServerDeployment(ctx context.Context, server runtime.ServerConfig, internalEndpoints []*schema.InternalEndpoint, opts deployOpts, s *serverRunState) error {
	srv := server.Server

	if server.Image.Repository == "" {
		return fnerrors.InternalError("kubernetes: no repository defined in image: %v", server.Image)
	}

	c, ok := constants[r.env.Purpose]
	if !ok {
		return fnerrors.InternalError("%s: no constants configured", r.env.Name)
	}

	kubepkg, err := srv.Env().LoadByName(ctx, kubeNode)
	if err != nil {
		return err
	}

	secCtx := applycorev1.SecurityContext()
	podSecCtx := applycorev1.PodSecurityContext()

	toparse := map[string]interface{}{
		"defaults/container.securitycontext.yaml": secCtx,
		"defaults/pod.podsecuritycontext.yaml":    podSecCtx,
	}

	for _, data := range kubepkg.PackageData {
		if obj, ok := toparse[data.Path]; ok {
			if err := yaml.Unmarshal(data.Contents, obj); err != nil {
				return fnerrors.InternalError("%s: failed to parse defaults: %w", data.Path, err)
			}
		}
	}

	if server.ReadOnlyFilesystem {
		secCtx = secCtx.WithReadOnlyRootFilesystem(true)
	}

	if server.RunAs != nil {
		userId, err := strconv.ParseInt(server.RunAs.UserID, 10, 64)
		if err != nil {
			return fnerrors.InternalError("expected server.RunAs to be an int64: %w", err)
		}

		podSecCtx = podSecCtx.WithRunAsUser(userId).WithRunAsNonRoot(true)
	}

	name := strings.ToLower(server.Server.Name()) // k8s doesn't accept uppercase names.
	containers := []string{name}
	container := applycorev1.Container().
		WithName(name).
		WithImage(server.Image.RepoAndDigest()).
		WithArgs(server.Args...).
		WithCommand(server.Command...).
		WithSecurityContext(secCtx)

	for _, internal := range internalEndpoints {
		for _, md := range internal.ServiceMetadata {
			if md.Kind == runtime.FnServiceLivez || md.Kind == runtime.FnServiceReadyz {
				probe, err := toProbe(internal, md)
				if err != nil {
					return err
				}

				probe = probe.WithPeriodSeconds(c.dashnessPeriod).WithFailureThreshold(c.failureThreshold).WithTimeoutSeconds(c.probeTimeout)

				switch md.Kind {
				case runtime.FnServiceLivez:
					container = container.WithLivenessProbe(probe.WithInitialDelaySeconds(c.livenessInitialDelay))
				case runtime.FnServiceReadyz:
					container = container.WithReadinessProbe(probe.WithInitialDelaySeconds(c.readinessInitialDelay))
				}
			}
		}
	}

	for _, kv := range server.Env {
		container = container.WithEnv(applycorev1.EnvVar().WithName(kv.Name).WithValue(kv.Value))
	}

	if server.WorkingDir != "" {
		container = container.WithWorkingDir(server.WorkingDir)
	}

	spec := applycorev1.PodSpec().
		WithSecurityContext(podSecCtx)

	labels := kubedef.MakeLabels(r.env, srv.Proto())
	annotations := kubedef.MakeAnnotations(srv.StackEntry())

	deploymentId := kubedef.MakeDeploymentId(srv.Proto())

	tmpl := applycorev1.PodTemplateSpec().
		WithAnnotations(annotations).
		WithLabels(labels)

	var initVolumeMounts []*applycorev1.VolumeMountApplyConfiguration
	initArgs := map[schema.PackageName][]string{}

	var serviceAccount string // May be specified by a SpecExtension.
	for _, input := range server.Extensions {
		specExt := &kubedef.SpecExtension{}
		containerExt := &kubedef.ContainerExtension{}
		initContainerExt := &kubedef.InitContainerExtension{}

		switch {
		case input.Impl.MessageIs(specExt):
			if err := input.Impl.UnmarshalTo(specExt); err != nil {
				return fnerrors.InternalError("failed to unmarshal SpecExtension: %w", err)
			}

			for _, vol := range specExt.Volume {
				k8svol, err := toK8sVol(vol)
				if err != nil {
					return err
				}
				spec = spec.WithVolumes(k8svol)
			}

			if len(specExt.Annotation) > 0 {
				m := map[string]string{}
				for _, annotation := range specExt.Annotation {
					m[annotation.Key] = annotation.Value
				}
				tmpl = tmpl.WithAnnotations(m)
			}

			if serviceAccount != "" && serviceAccount != specExt.ServiceAccount {
				return fnerrors.UserError(server.Server.Location, "incompatible service accounts defined, %q vs %q", serviceAccount, specExt.ServiceAccount)
			}

			serviceAccount = specExt.ServiceAccount

		case input.Impl.MessageIs(containerExt):
			if err := input.Impl.UnmarshalTo(containerExt); err != nil {
				return fnerrors.InternalError("failed to unmarshal ContainerExtension: %w", err)
			}

			for _, mount := range containerExt.VolumeMount {
				volumeMount := applycorev1.VolumeMount().
					WithName(mount.Name).
					WithReadOnly(mount.ReadOnly).
					WithMountPath(mount.MountPath)
				container = container.WithVolumeMounts(volumeMount)
				if mount.MountOnInit {
					// Volume mounts may declare to be available also during server initialization.
					// E.g. Initializing the schema of a data store requires early access to server secrets.
					// The volume mount provider has full control over whether the volume is available.
					initVolumeMounts = append(initVolumeMounts, volumeMount)
				}
			}

			// XXX O(n^2)
			for _, env := range containerExt.Env {
				if currentValue, found := getEnv(container, env.Name); found && currentValue != env.Value {
					return fnerrors.UserError(server.Server.Location, "env variable '%s' is already set to '%s' but would be overwritten to '%s' by container extension", env.Name, currentValue, env.Value)
				}
				container = container.WithEnv(applycorev1.EnvVar().WithName(env.Name).WithValue(env.Value))
			}

			if containerExt.Args != nil {
				container = container.WithArgs(containerExt.Args...)
			} else {
				// Deprecated path.
				for _, arg := range containerExt.ArgTuple {
					if currentValue, found := getArg(container, arg.Name); found && currentValue != arg.Value {
						return fnerrors.UserError(server.Server.Location, "argument '%s' is already set to '%s' but would be overwritten to '%s' by container extension", arg.Name, currentValue, arg.Value)
					}
					container = container.WithArgs(fmt.Sprintf("--%s=%s", arg.Name, arg.Value))
				}
			}

			// Deprecated path.
			for _, initContainer := range containerExt.InitContainer {
				pkg := schema.PackageName(initContainer.PackageName)
				initArgs[pkg] = append(initArgs[pkg], initContainer.Arg...)
			}

		case input.Impl.MessageIs(initContainerExt):
			if err := input.Impl.UnmarshalTo(initContainerExt); err != nil {
				return fnerrors.InternalError("failed to unmarshal InitContainerExtension: %w", err)
			}

			pkg := schema.PackageName(initContainerExt.PackageName)
			initArgs[pkg] = append(initArgs[pkg], initContainerExt.Args...)

		default:
			return fnerrors.InternalError("unused startup input: %s", input.Impl.GetTypeUrl())
		}
	}

	for _, rs := range srv.Proto().RequiredStorage {
		if rs.Owner == "" {
			return fnerrors.UserError(server.Server.Location, "requiredstorage owner is not set")
		}

		if rs.ByteCount >= math.MaxInt64 {
			return fnerrors.UserError(server.Server.Location, "requiredstorage value too high (maximum is %d)", math.MaxInt64)
		}

		container = container.WithVolumeMounts(
			applycorev1.VolumeMount().
				WithName(makeStorageVolumeName(rs)).
				WithMountPath(rs.MountPath))
		spec = spec.WithVolumes(applycorev1.Volume().
			WithName(makeStorageVolumeName(rs)).
			WithPersistentVolumeClaim(
				applycorev1.PersistentVolumeClaimVolumeSource().
					WithClaimName(rs.PersistentId)))

		s.declarations = append(s.declarations, kubedef.Apply{
			Description: fmt.Sprintf("Persistent storage for %s", rs.Owner),
			Resource:    "persistentvolumeclaims",
			Namespace:   r.ns(),
			Name:        rs.PersistentId,
			Body: applycorev1.PersistentVolumeClaim(rs.PersistentId, r.ns()).
				WithSpec(applycorev1.PersistentVolumeClaimSpec().
					WithAccessModes(corev1.ReadWriteOnce).
					WithResources(applycorev1.ResourceRequirements().WithRequests(corev1.ResourceList{
						corev1.ResourceStorage: *resource.NewScaledQuantity(int64(rs.ByteCount), resource.Scale(1)),
					}))),
		})
	}

	for _, init := range server.Inits {
		name := fmt.Sprintf("init-%v", labelName(init.Command))
		for _, c := range containers {
			if name == c {
				return fnerrors.UserError(server.Server.Location, "duplicate init container name: %s", name)
			}
		}
		containers = append(containers, name)

		spec.WithInitContainers(
			applycorev1.Container().
				WithName(name).
				WithImage(init.Image.RepoAndDigest()).
				WithArgs(append(init.Args, initArgs[init.PackageName]...)...).
				WithCommand(init.Command...).
				WithVolumeMounts(initVolumeMounts...))
	}

	spec = spec.
		WithContainers(container).
		WithAutomountServiceAccountToken(serviceAccount != "")

	if serviceAccount != "" {
		spec = spec.WithServiceAccountName(serviceAccount)
	}

	tmpl = tmpl.WithSpec(spec)

	// Only mutate `annotations` after all other uses above.
	if server.ConfigImage != nil {
		annotations[kubedef.K8sConfigImage] = server.ConfigImage.RepoAndDigest()
	}

	if server.Server.IsStateful() {
		s.declarations = append(s.declarations, kubedef.Apply{
			Description: "Server StatefulSet",
			Resource:    "statefulsets",
			Namespace:   r.ns(),
			Name:        deploymentId,
			Body: appsv1.
				StatefulSet(deploymentId, r.ns()).
				WithAnnotations(annotations).
				WithLabels(labels).
				WithSpec(appsv1.StatefulSetSpec().
					WithReplicas(1).
					WithTemplate(tmpl).
					WithSelector(applymetav1.LabelSelector().WithMatchLabels(kubedef.SelectById(srv.Proto())))),
		})
	} else {
		s.declarations = append(s.declarations, kubedef.Apply{
			Description: "Server Deployment",
			Resource:    "deployments",
			Namespace:   r.ns(),
			Name:        deploymentId,
			Body: appsv1.
				Deployment(deploymentId, r.ns()).
				WithAnnotations(annotations).
				WithLabels(labels).
				WithSpec(appsv1.DeploymentSpec().
					WithReplicas(1).
					WithTemplate(tmpl).
					WithSelector(applymetav1.LabelSelector().WithMatchLabels(kubedef.SelectById(srv.Proto())))),
		})
	}

	return nil
}

func makeStorageVolumeName(rs *schema.RequiredStorage) string {
	h := sha256.New()
	fmt.Fprint(h, rs.Owner)
	return "rs-" + hex.EncodeToString(h.Sum(nil))[:8]
}

func (r boundEnv) deployEndpoint(ctx context.Context, server runtime.ServerConfig, endpoint *schema.Endpoint, s *serverRunState) error {
	t := server.Server

	serviceSpec := applycorev1.ServiceSpec().WithSelector(kubedef.SelectById(t.Proto()))

	port := endpoint.Port
	if port != nil {
		serviceSpec = serviceSpec.WithPorts(applycorev1.ServicePort().
			WithProtocol(corev1.ProtocolTCP).WithName(port.Name).WithPort(port.ContainerPort))

		serviceAnnotations, err := kubedef.MakeServiceAnnotations(t.Proto(), endpoint)
		if err != nil {
			return err
		}

		s.declarations = append(s.declarations, kubedef.Apply{
			Description: fmt.Sprintf("Service %s", endpoint.ServiceName),
			Resource:    "services",
			Namespace:   r.ns(),
			Name:        endpoint.AllocatedName,
			Body: applycorev1.
				Service(endpoint.AllocatedName, r.ns()).
				WithLabels(kubedef.MakeServiceLabels(r.env, t.Proto(), endpoint)).
				WithAnnotations(serviceAnnotations).
				WithSpec(serviceSpec),
		})
	}

	return nil
}

func toK8sVol(vol *kubedef.SpecExtension_Volume) (*applycorev1.VolumeApplyConfiguration, error) {
	v := applycorev1.Volume().WithName(vol.Name)
	switch x := vol.VolumeType.(type) {
	case *kubedef.SpecExtension_Volume_Secret_:
		return v.WithSecret(applycorev1.SecretVolumeSource().WithSecretName(x.Secret.SecretName)), nil
	case *kubedef.SpecExtension_Volume_ConfigMap_:
		vol := applycorev1.ConfigMapVolumeSource().WithName(x.ConfigMap.Name)
		for _, it := range x.ConfigMap.Item {
			vol = vol.WithItems(applycorev1.KeyToPath().WithKey(it.Key).WithPath(it.Path))
		}
		return v.WithConfigMap(vol), nil
	default:
		return nil, fnerrors.InternalError("don't know how to instantiate a k8s volume from %v", vol)
	}
}
