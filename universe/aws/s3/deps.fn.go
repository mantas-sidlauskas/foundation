// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `ns generate`.

package s3

import (
	"context"
	fncore "namespacelabs.dev/foundation/std/core"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/universe/aws/client"
)

// Dependencies that are instantiated once for the lifetime of the extension.
type ExtensionDeps struct {
	ClientFactory  client.ClientFactory
	ReadinessCheck core.Check
}

type _checkProvideBucket func(context.Context, *BucketConfig, ExtensionDeps) (*Bucket, error)

var _ _checkProvideBucket = ProvideBucket

var (
	Package__eoj2dq = &core.Package{
		PackageName: "namespacelabs.dev/foundation/universe/aws/s3",
	}

	Provider__eoj2dq = core.Provider{
		Package:     Package__eoj2dq,
		Instantiate: makeDeps__eoj2dq,
	}
)

func makeDeps__eoj2dq(ctx context.Context, di core.Dependencies) (_ interface{}, err error) {
	var deps ExtensionDeps

	if err := di.Instantiate(ctx, client.Provider__hva50k, func(ctx context.Context, v interface{}) (err error) {
		if deps.ClientFactory, err = client.ProvideClientFactory(ctx, nil, v.(client.ExtensionDeps)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if deps.ReadinessCheck, err = fncore.ProvideReadinessCheck(ctx, nil); err != nil {
		return nil, err
	}

	return deps, nil
}
