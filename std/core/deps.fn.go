// This file was automatically generated by Namespace.
// DO NOT EDIT. To update, re-run `ns generate`.

package core

import (
	"context"
	"namespacelabs.dev/foundation/std/core/types"
	"namespacelabs.dev/foundation/std/go/core"
)

type _checkProvideDebugHandler func(context.Context, *DebugHandlerArgs) (core.DebugHandler, error)

var _ _checkProvideDebugHandler = ProvideDebugHandler

type _checkProvideLivenessCheck func(context.Context, *LivenessCheckArgs) (core.Check, error)

var _ _checkProvideLivenessCheck = ProvideLivenessCheck

type _checkProvideReadinessCheck func(context.Context, *ReadinessCheckArgs) (core.Check, error)

var _ _checkProvideReadinessCheck = ProvideReadinessCheck

type _checkProvideServerInfo func(context.Context, *ServerInfoArgs) (*types.ServerInfo, error)

var _ _checkProvideServerInfo = ProvideServerInfo

var (
	Package__9ae6je = &core.Package{
		PackageName: "namespacelabs.dev/foundation/std/core",
	}
)