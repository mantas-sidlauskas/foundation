// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package cuefrontend

import (
	"context"
	"errors"

	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/fnfs/memfs"
	"namespacelabs.dev/foundation/internal/frontend/fncue"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
)

type impl struct {
	loader  workspace.EarlyPackageLoader
	evalctx *fncue.EvalCtx
}

func NewFrontend(pl workspace.EarlyPackageLoader) workspace.Frontend {
	return impl{
		loader:  pl,
		evalctx: fncue.NewEvalCtx(workspaceLoader{pl}),
	}
}

func (ft impl) ParsePackage(ctx context.Context, loc workspace.Location, opts workspace.LoadPackageOpts) (*workspace.Package, error) {
	partial, err := parsePackage(ctx, ft.evalctx, ft.loader, loc)
	if err != nil {
		return nil, err
	}

	v := &partial.CueV

	parsed := &workspace.Package{
		Location: loc,
		Parsed:   phase1plan{Value: v, Left: partial.Left},
	}

	var count int
	if extension := v.LookupPath("extension"); extension.Exists() {
		if err := parseCueNode(ctx, ft.loader, loc, schema.Node_EXTENSION, v, extension, parsed, opts); err != nil {
			return nil, fnerrors.Wrapf(loc, err, "parsing extension")
		}
		count++
	}

	if service := v.LookupPath("service"); service.Exists() {
		if err := parseCueNode(ctx, ft.loader, loc, schema.Node_SERVICE, v, service, parsed, opts); err != nil {
			return nil, fnerrors.Wrapf(loc, err, "parsing service")
		}
		count++
	}

	if server := v.LookupPath("server"); server.Exists() {
		parsedSrv, err := parseCueServer(ctx, ft.loader, loc, v, server, parsed, opts)
		if err != nil {
			return nil, fnerrors.Wrapf(loc, err, "parsing server")
		}
		parsed.Server = parsedSrv
		count++
	}

	if binary := v.LookupPath("binary"); binary.Exists() {
		parsedBinary, err := parseCueBinary(ctx, loc, v, binary)
		if err != nil {
			return nil, fnerrors.Wrapf(loc, err, "parsing binary")
		}
		parsed.Binary = parsedBinary
		count++
	}

	if test := v.LookupPath("test"); test.Exists() {
		parsedTest, err := parseCueTest(ctx, loc, v, test)
		if err != nil {
			return nil, fnerrors.Wrapf(loc, err, "parsing test")
		}
		parsed.Test = parsedTest
		count++
	}

	if count > 1 {
		return nil, errors.New("package must only define one of: server, service, extension, binary or test")
	}

	return parsed, nil
}

func (ft impl) HasPackage(ctx context.Context, pkg schema.PackageName) (bool, error) {
	firstPass, err := ft.evalctx.Eval(ctx, pkg.String())
	if err != nil {
		return false, err
	}

	var topLevels = []string{"service", "server", "extension", "test", "binary"}
	for _, topLevel := range topLevels {
		if firstPass.LookupPath(topLevel).Exists() {
			return true, nil
		}
	}

	return false, nil
}

type workspaceLoader struct {
	pl workspace.EarlyPackageLoader
}

func (wl workspaceLoader) SnapshotDir(ctx context.Context, pkgname schema.PackageName, opts memfs.SnapshotOpts) (fnfs.Location, error) {
	loc, err := wl.pl.Resolve(ctx, pkgname)
	if err != nil {
		return fnfs.Location{}, err
	}

	w, err := wl.pl.WorkspaceOf(ctx, loc.Module)
	if err != nil {
		return fnfs.Location{}, err
	}

	fsys, err := w.SnapshotDir(loc.Rel(), opts)
	if err != nil {
		return fnfs.Location{}, err
	}

	return fnfs.Location{
		ModuleName: loc.Module.ModuleName(),
		RelPath:    loc.Rel(),
		FS:         fsys,
	}, nil
}