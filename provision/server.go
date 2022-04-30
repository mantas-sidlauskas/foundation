// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package provision

import (
	"context"
	"fmt"
	"strings"

	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/frontend"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
)

// Represents a server bound to an environment.
type Server struct {
	Location workspace.Location
	Package  *workspace.Package

	env   ServerEnv            // The environment this server instance is bound to.
	entry *schema.Stack_Entry  // The stack entry, i.e. all of the server's dependencies.
	deps  []*workspace.Package // List of parsed deps.
}

type ServerEnv interface {
	workspace.WorkspaceEnvironment
	workspace.SealedPackages
}

func (t Server) Module() *workspace.Module       { return t.Location.Module }
func (t Server) Env() ServerEnv                  { return t.env }
func (t Server) PackageName() schema.PackageName { return t.Location.PackageName }
func (t Server) StackEntry() *schema.Stack_Entry { return t.entry }
func (t Server) Proto() *schema.Server           { return t.entry.Server }
func (t Server) Name() string                    { return t.entry.Server.Name }
func (t Server) Framework() schema.Framework     { return t.entry.Server.Framework }
func (t Server) IsStateful() bool                { return t.entry.Server.IsStateful }
func (t Server) Deps() []*workspace.Package      { return t.deps }

func (t Server) GetDep(pkg schema.PackageName) *workspace.Package {
	for _, d := range t.deps {
		if d.PackageName() == pkg {
			return d
		}
	}
	return nil
}

func makeServer(ctx context.Context, loader workspace.Packages, env *schema.Environment, pkgname schema.PackageName, bind func() ServerEnv) (Server, error) {
	sealed, err := workspace.Seal(ctx, loader, pkgname, &workspace.SealHelper{
		AdditionalServerDeps: func() ([]schema.PackageName, error) {
			return []schema.PackageName{
				schema.PackageName(fmt.Sprintf("namespacelabs.dev/foundation/std/runtime/%s", strings.ToLower(env.Runtime))),
			}, nil
		},
	})
	if err != nil {
		return Server{}, err
	}

	t := Server{
		Location: sealed.ParsedPackage.Location,
		env:      bind(),
	}

	t.Package = sealed.ParsedPackage
	t.entry = sealed.Proto
	t.deps = sealed.Deps

	pdata, err := t.Package.Parsed.EvalProvision(ctx, t.Env(), frontend.ProvisionInputs{
		Workspace:      t.Module().Workspace,
		ServerLocation: t.Location,
	})
	if err != nil {
		return Server{}, fnerrors.Wrap(t.Location, err)
	}

	if len(pdata.DeclaredStack) > 0 {
		return Server{}, fnerrors.UserError(t.Location, "servers can't add servers to the stack")
	}

	if len(pdata.Sidecars) > 0 {
		return Server{}, fnerrors.UserError(t.Location, "servers can't define sidecar containers")
	}

	if len(pdata.Inits) > 0 {
		return Server{}, fnerrors.UserError(t.Location, "servers can't define init containers")
	}

	if pdata.Provisioning != nil {
		return Server{}, fnerrors.UserError(t.Location, "servers can't specify a custom provisioning phase")
	}

	if len(pdata.Details) > 0 {
		return Server{}, fnerrors.UserError(t.Location, "servers can't specify a custom details")
	}

	t.entry.ServerNaming = pdata.Naming

	return t, nil
}

func ServerSchemas(servers []Server) []*schema.Server {
	var s []*schema.Server
	for _, srv := range servers {
		s = append(s, srv.Proto())
	}
	return s
}
