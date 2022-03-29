// This file was automatically generated.
package grpcservice

import (
	"context"

	"namespacelabs.dev/foundation/std/go/grpc/server"
	"namespacelabs.dev/foundation/std/testdata/go/datastore"
)

type ServiceDeps struct {
	Main *datastore.DB
}

// Verify that WireService is present and has the appropriate type.
type checkWireService func(context.Context, *server.Grpc, ServiceDeps)

var _ checkWireService = WireService
