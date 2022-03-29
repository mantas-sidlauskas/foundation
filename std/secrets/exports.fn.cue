// This file is automatically generated.
package secrets

#Exports: {
	Secret: {
		packageName: "namespacelabs.dev/foundation/std/secrets"
		type:        "Secret"
		typeDefinition: {
			"typename": "ns.Secret"
			"source": [
				"provider.proto",
			]
		}
		with: {
			name?: string
			provision?: [...("PROVISION_UNSPECIFIED" | "PROVISION_INLINE" | "PROVISION_AS_FILE")]
		}
	}
}
