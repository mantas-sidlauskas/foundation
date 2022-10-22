// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/build/buildkit"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/compute/cache"
	"namespacelabs.dev/foundation/internal/parsing/devhost"
	"namespacelabs.dev/foundation/internal/parsing/module"
	"namespacelabs.dev/foundation/internal/unprepare"
	"namespacelabs.dev/foundation/std/tasks"
)

func NewUnprepareCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unprepare",
		Aliases: []string{"clean"},
		Short:   "Reverts the local workspace to zero state by deleting all Namespace-managed containers and caches.",
		Args:    cobra.NoArgs,

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			root, err := module.FindRoot(ctx, ".")
			if err != nil {
				return err
			}

			// Remove k3d cluster(s) and registry(ies).
			if err := unprepare.UnprepareK3d(ctx); err != nil {
				return err
			}

			// Stop and remove the builtkit daemon container.
			if err := buildkit.RemoveBuildkitd(ctx); err != nil {
				return err
			}

			// Remove the devhost configuration.
			if err := tasks.Action("delete.devhost").Run(ctx, func(ctx context.Context) error {
				return devhost.Delete(ctx, root)
			}); err != nil {
				return err
			}

			// Prune cached build artifacts and command history artifacts.
			if err := cache.Prune(ctx); err != nil {
				return err
			}

			return nil
		}),
	}

	return cmd
}
