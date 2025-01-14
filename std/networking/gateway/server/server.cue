import (
	"namespacelabs.dev/foundation/std/fn"
)

server: fn.#OpaqueServer & {
	id:   "sun4qtee50l61888bdj0"
	name: "gateway"

	binary: image: "envoyproxy/envoy:v1.22.0@sha256:478044b54936608dd3115c89ea9fe5be670f1e78d4436927c096b4bc06eeedeb"

	service: {
		"admin": {
			label:         "Envoy (admin)"
			containerPort: 19000
			metadata: protocol: "http"
		}

		// Must be consistent with controller's configuration.
		"grpc-http-transcoder": {
			containerPort: 10000
			metadata: protocol: "http"
			internal: true // Not used for development.
		}
	}
}

configure: fn.#Configure & {
	with: binary: "namespacelabs.dev/foundation/std/networking/gateway/server/configure"

	sidecar: controller: {
		binary: "namespacelabs.dev/foundation/std/networking/gateway/controller"
	}
}
