// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package debug

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/prototext"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/parsing/module"
	"namespacelabs.dev/foundation/internal/runtime/kubernetes"
	"namespacelabs.dev/foundation/std/cfg"
)

func newKubernetesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kubernetes",
	}

	envBound := "dev"
	systemInfo := &cobra.Command{
		Use:  "system-info",
		Args: cobra.NoArgs,

		RunE: fncobra.RunE(func(ctx context.Context, args []string) error {
			root, err := module.FindRoot(ctx, ".")
			if err != nil {
				return err
			}

			env, err := cfg.LoadContext(root, envBound)
			if err != nil {
				return err
			}

			k, err := kubernetes.ConnectToCluster(ctx, env.Configuration())
			if err != nil {
				return err
			}

			sysInfo, err := k.SystemInfo(ctx)
			if err != nil {
				return err
			}

			fmt.Fprintln(console.Stdout(ctx), prototext.Format(sysInfo))
			return nil
		}),
	}

	systemInfo.Flags().StringVar(&envBound, "env", envBound, "If specified, produce a env-bound sealed schema.")

	cmd.AddCommand(systemInfo)

	return cmd
}
