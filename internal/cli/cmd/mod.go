// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/parsing"
	"namespacelabs.dev/foundation/internal/parsing/module"
)

func NewModCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mod",
		Short: "Module related operations (e.g. download, get, tidy).",
	}

	cmd.AddCommand(NewTidyCmd())
	cmd.AddCommand(newModDownloadCmd())
	cmd.AddCommand(newModGetCmd())

	return cmd
}

func newModDownloadCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Downloads all referenced modules.",

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			root, err := module.FindRootWithArgs(ctx, ".", parsing.ModuleAtArgs{SkipAPIRequirements: true})
			if err != nil {
				return err
			}

			for _, dep := range root.Workspace().Proto().Dep {
				mod, err := parsing.DownloadModule(ctx, dep, force)
				if err != nil {
					return err
				}

				fmt.Fprintf(console.Stdout(ctx), "Downloaded %s: %s\n", mod.ModuleName, mod.Version)
			}

			return nil
		}),
	}

	cmd.Flags().BoolVar(&force, "force", force, "Download a module even if it already exists locally.")

	return cmd
}

func newModGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <module-uri>",
		Short: "Gets the latest version of the specified module.",
		Args:  cobra.ExactArgs(1),

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			root, err := module.FindRootWithArgs(ctx, ".", parsing.ModuleAtArgs{SkipAPIRequirements: true})
			if err != nil {
				return err
			}

			dep, err := parsing.ResolveModuleVersion(ctx, args[0])
			if err != nil {
				return err
			}

			if _, err := parsing.DownloadModule(ctx, dep, false); err != nil {
				return err
			}

			newData := root.EditableWorkspace().WithSetDependency(dep)
			if newData != nil {
				return rewriteWorkspace(ctx, root, newData)
			}

			return nil
		}),
	}

	return cmd
}
