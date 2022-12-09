// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package runtime

import (
	"context"
	"fmt"

	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/cfg"
)

const (
	LocalBaseDomain = "nslocal.host"
	CloudBaseDomain = "nscloud.dev"
)

var (
	NamingNoTLS = false // Set to true in CI.

	WorkInProgressUseShortAlias = false
)

func ComputeNaming(ctx context.Context, ws string, env cfg.Context, cluster Planner, source *schema.Naming) (*schema.ComputedNaming, error) {
	result, err := computeNaming(ctx, ws, env, cluster, source)
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(console.Debug(ctx), "computed naming: %+v\n", result)
	return result, nil
}

func computeNaming(ctx context.Context, workspace string, env cfg.Context, cluster Planner, source *schema.Naming) (*schema.ComputedNaming, error) {
	naming, err := computeInnerNaming(ctx, env, cluster, source)
	if err != nil {
		return nil, err
	}

	naming.MainModuleName = workspace
	naming.UseShortAlias = naming.GetUseShortAlias() || WorkInProgressUseShortAlias

	return naming, nil
}

func computeInnerNaming(ctx context.Context, rootenv cfg.Context, cluster Planner, source *schema.Naming) (*schema.ComputedNaming, error) {
	base, err := cluster.ComputeBaseNaming(source)
	if err != nil {
		return nil, err
	}

	if base != nil {
		return base, nil
	}

	env := rootenv.Environment()

	if env.Purpose != schema.Environment_PRODUCTION {
		return &schema.ComputedNaming{
			Source:     source,
			BaseDomain: LocalBaseDomain,
			Managed:    schema.Domain_LOCAL_MANAGED,
		}, nil
	}

	if !source.GetEnableNamespaceManaged() {
		return &schema.ComputedNaming{}, nil
	}

	org := source.GetWithOrg()
	if org == "" {
		return &schema.ComputedNaming{}, nil
	}

	return &schema.ComputedNaming{
		Source:     source,
		BaseDomain: fmt.Sprintf("%s.%s", org, CloudBaseDomain),
		Managed:    schema.Domain_CLOUD_MANAGED,
	}, nil
}

func allocateName(ctx context.Context, srv Deployable, opts fnapi.AllocateOpts) (*schema.Certificate, error) {
	opts.NoTLS = NamingNoTLS
	opts.Scope = srv.GetPackageRef().AsPackageName()

	nr, err := fnapi.AllocateName(ctx, opts)
	if err != nil {
		return nil, err
	}

	return certFromResource(nr), nil
}

func certFromResource(res *fnapi.NameResource) *schema.Certificate {
	if res.Certificate.PrivateKey != nil && res.Certificate.CertificateBundle != nil {
		return &schema.Certificate{
			PrivateKey:        res.Certificate.PrivateKey,
			CertificateBundle: res.Certificate.CertificateBundle,
		}
	}

	return nil
}
