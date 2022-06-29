// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/morikuni/aec"
	"github.com/spf13/cobra"
	"namespacelabs.dev/foundation/devworkflow"
	"namespacelabs.dev/foundation/devworkflow/keyboard"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/reverseproxy"
	"namespacelabs.dev/foundation/languages/web"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
	"namespacelabs.dev/foundation/workspace/compute"
	"namespacelabs.dev/foundation/workspace/module"
	"namespacelabs.dev/foundation/workspace/tasks"
)

func NewDevCmd() *cobra.Command {
	var (
		envRef       = "dev"
		servingAddr  string
		devWebServer = false
	)

	cmd := &cobra.Command{
		Use:   "dev",
		Short: "Starts a development session, continuously building and deploying a server.",
		Args:  cobra.MinimumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, sink := tasks.WithStatefulSink(cmd.Context())

			ctxWithCancel, cancel := fncobra.WithSigIntCancel(ctx)
			defer cancel()

			return compute.Do(ctxWithCancel, func(ctx context.Context) error {
				root, err := module.FindRoot(ctx, ".")
				if err != nil {
					return err
				}

				lis, err := startListener(servingAddr)
				if err != nil {
					return fnerrors.InternalError("Failed to start listener at %q: %w", servingAddr, err)
				}

				defer lis.Close()

				pl := workspace.NewPackageLoader(root)

				var serverPackages []string
				var serverProtos []*schema.Server
				for _, p := range args {
					parsed, err := pl.LoadByName(ctx, root.RelPackage(p).AsPackageName())
					if err != nil {
						return err
					}

					if parsed.Server == nil {
						return fnerrors.UserError(parsed.Location, "`ns dev` works exclusively with servers (for now)")
					}

					serverPackages = append(serverPackages, parsed.PackageName().String())
					serverProtos = append(serverProtos, parsed.Server)
				}

				localHost := lis.Addr().(*net.TCPAddr).IP.String()

				updateWebUISticky(ctx, "preparing")

				sesh, err := devworkflow.NewSession(ctx, sink, localHost)
				if err != nil {
					return err
				}

				console.SetIdleLabel(ctx, "waiting for workspace changes")

				// Kick off the dev workflow.
				sesh.DeferRequest(&devworkflow.DevWorkflowRequest{
					Type: &devworkflow.DevWorkflowRequest_SetWorkspace_{
						SetWorkspace: &devworkflow.DevWorkflowRequest_SetWorkspace{
							AbsRoot:           root.Abs(),
							PackageName:       serverPackages[0],
							AdditionalServers: serverPackages[1:],
							EnvName:           envRef,
						},
					},
				})

				r := mux.NewRouter()
				fncobra.RegisterPprof(r)
				devworkflow.RegisterEndpoints(sesh, r)

				if err := keyboard.StartHandler(ctx, sesh, root, serverProtos, cancel); err != nil {
					return err
				}

				if devWebServer {
					localPort := lis.Addr().(*net.TCPAddr).Port
					webPort := localPort + 1
					proxyTarget, err := web.StartDevServer(ctx, root, devworkflow.WebPackage, localPort, webPort)
					if err != nil {
						return err
					}
					r.PathPrefix("/").Handler(reverseproxy.Make(proxyTarget, reverseproxy.DefaultLocalProxy()))
				} else {
					mux, err := devworkflow.PrebuiltWebUI(ctx)
					if err != nil {
						return err
					}

					r.PathPrefix("/").Handler(mux)
				}

				srv := &http.Server{
					Handler:      r,
					Addr:         servingAddr,
					WriteTimeout: 15 * time.Second,
					ReadTimeout:  15 * time.Second,
					BaseContext:  func(l net.Listener) context.Context { return ctx },
				}

				sesh.RunInContext(func(ctx context.Context) error {
					// On cancelation, i.e. Ctrl-C, ask the server to shutdown. This will lead to the next go-routine below, actually returns.
					<-ctx.Done()

					ctxT, cancelT := context.WithTimeout(ctx, 1*time.Second)
					defer cancelT()

					return srv.Shutdown(ctxT)
				})

				sesh.RunInContext(func(ctx context.Context) error {
					updateWebUISticky(ctx, "running at: http://%s", lis.Addr())
					return srv.Serve(lis)
				})

				return sesh.Run(ctx)
			})
		},
	}

	cmd.Flags().StringVarP(&servingAddr, "listen", "H", "", "Listen on the specified address.")
	cmd.Flags().StringVar(&envRef, "env", envRef, "The environment to provision (as defined in the workspace).")
	cmd.Flags().BoolVar(&devWebServer, "devweb", devWebServer, "Whether to start a development web frontend.")
	cmd.Flags().BoolVar(&devworkflow.AlsoOutputBuildToStderr, "alsooutputtostderr", devworkflow.AlsoOutputBuildToStderr, "Also send build output to stderr.")

	return cmd
}

func updateWebUISticky(ctx context.Context, format string, args ...any) {
	console.SetStickyContent(ctx, "webui", fmt.Sprintf(" %s: web ui %s", aec.Bold.Apply("Namespace"), fmt.Sprintf(format, args...)))
}

func startListener(specified string) (net.Listener, error) {
	const defaultHostname = "127.0.0.1"
	const defaultStartingPort = 4001

	if specified != "" {
		return net.Listen("tcp", specified)
	}

	for port := defaultStartingPort; ; port++ {
		l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", defaultHostname, port))
		if err != nil {
			var errno syscall.Errno
			if errors.As(err, &errno) {
				if errno == syscall.EADDRINUSE {
					continue
				}
			}
			return nil, err
		}

		return l, nil
	}
}
