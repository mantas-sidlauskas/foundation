// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `ns generate`.

package middleware

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
)

type _checkProvideMiddleware func(context.Context, *MiddlewareRegistration) (Middleware, error)

var _ _checkProvideMiddleware = ProvideMiddleware

var (
	Package__c3bggl = &core.Package{
		PackageName: "namespacelabs.dev/foundation/std/go/http/middleware",
	}
)
