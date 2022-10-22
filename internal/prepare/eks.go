// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package prepare

import (
	"context"

	"namespacelabs.dev/foundation/internal/compute"
	"namespacelabs.dev/foundation/internal/parsing/devhost"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/cfg"
	"namespacelabs.dev/foundation/std/tasks"
	"namespacelabs.dev/foundation/universe/aws/configuration/eks"
)

func PrepareEksCluster(env cfg.Context, clusterName string) compute.Computable[[]*schema.DevHost_ConfigureEnvironment] {
	return compute.Map(
		tasks.Action("prepare.eks-cluster-config").HumanReadablef("Prepare the EKS cluster configuration"),
		compute.Inputs().Str("clusterName", clusterName).Proto("env", env.Environment()),
		compute.Output{NotCacheable: true},
		func(ctx context.Context, _ compute.Resolved) ([]*schema.DevHost_ConfigureEnvironment, error) {
			c, err := devhost.MakeConfiguration(&eks.Cluster{Name: clusterName})
			if err != nil {
				return nil, err
			}
			c.Name = env.Environment().Name
			return []*schema.DevHost_ConfigureEnvironment{c}, nil
		})
}
