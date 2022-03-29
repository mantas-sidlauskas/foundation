import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/fn:inputs"
	"namespacelabs.dev/foundation/universe/db/postgres/creds"
	"namespacelabs.dev/foundation/std/go/core"
)

$providerProto: inputs.#Proto & {
	sources: [
		"provider.proto",
		"../database.proto",
	]
}

extension: fn.#Extension & {
	import: [
		"namespacelabs.dev/foundation/universe/db/postgres",
	]

	instantiate: {
		"creds":        creds.#Exports.Creds
		readinessCheck: core.#Exports.ReadinessCheck
	}

	provides: {
		Database: {
			input: $providerProto.types.Database

			availableIn: {
				go: {
					package: "github.com/jackc/pgx/v4/pgxpool"
					type:    "*Pool"
				}
			}
		}
	}
}

configure: fn.#Configure & {
	with: binary: "namespacelabs.dev/foundation/universe/db/postgres/opaque/tool"
}
