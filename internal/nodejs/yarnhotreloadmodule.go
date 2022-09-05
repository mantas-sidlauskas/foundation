// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package nodejs

import (
	"context"
	"fmt"

	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs/workspace/wsremote"
	"namespacelabs.dev/foundation/internal/wscontents"
	"namespacelabs.dev/foundation/std/pkggraph"
	"namespacelabs.dev/foundation/workspace/compute"
)

const yarnLockFn = "yarn.lock"

type YarnHotReloadModule struct {
	Module *pkggraph.Module
	Sink   wsremote.Sink
}

func (w YarnHotReloadModule) ModuleName() string { return w.Module.ModuleName() }
func (w YarnHotReloadModule) Abs() string        { return w.Module.Abs() }
func (w YarnHotReloadModule) VersionedFS(rel string, observeChanges bool) compute.Computable[wscontents.Versioned] {
	if observeChanges {
		return wsremote.ObserveAndPush(w.Module.Abs(), rel, observerSink{w.Sink})
	}

	return w.Module.VersionedFS(rel, observeChanges)
}

type observerSink struct {
	sink wsremote.Sink
}

func (obs observerSink) Deposit(ctx context.Context, events []*wscontents.FileEvent) error {
	for _, ev := range events {
		if ev.Path == yarnLockFn {
			return fnerrors.ExpectedError(fmt.Sprintf("%s changed, triggering a rebuild", yarnLockFn))
		}
	}

	return obs.sink.Deposit(ctx, events)
}
