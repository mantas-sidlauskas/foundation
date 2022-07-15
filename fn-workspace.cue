module: "namespacelabs.dev/foundation"
requirements: {
	api: 37
}
prebuilts: {
	digest: {
		"namespacelabs.dev/foundation/cmd/ns":                                                         "sha256:94a3caac122be692d78085d15e22db5c25c9d53e83fa8ab5f5530ca2b0d4a55d"
		"namespacelabs.dev/foundation/cmd/nspipelines":                                                "sha256:c747350f8ad4af0167d02f238a8de434f8619d32177696522dbe8f36e3d79db4"
		"namespacelabs.dev/foundation/devworkflow/web":                                                "sha256:2b17750f189a5b812d713fa57539cd28764743dc1afd592ad6e770455628adc3"
		"namespacelabs.dev/foundation/internal/sdk/buf/image/prebuilt":                                "sha256:f2cf5502f9b6afc27f73386fbbfd31954434968b048723f37e22996228972ee3"
		"namespacelabs.dev/foundation/std/dev/controller":                                             "sha256:158c0d655240ac1d4c2360f7c1399433aef6e508eece3e048e23aa98ce2441c1"
		"namespacelabs.dev/foundation/std/development/filesync/controller":                            "sha256:2573188b741b370989262a91b45875e8a340942213a3c0def60b08287aca1848"
		"namespacelabs.dev/foundation/std/grpc/httptranscoding/configure":                             "sha256:9fcefb7e686ebb9f442dec5bdf65f352c2d5a004fc5408578aebfc9ad06e8bdf"
		"namespacelabs.dev/foundation/std/monitoring/grafana/tool":                                    "sha256:9e4fbae952218b45c004012b070a56075c2ac52adfd915a962e34a945f3fd78b"
		"namespacelabs.dev/foundation/std/monitoring/prometheus/tool":                                 "sha256:6fd4b1b41b8fc4e5f51bc111312beaab9fef9bfddb08090c41df1c50aa532e36"
		"namespacelabs.dev/foundation/std/networking/gateway/controller":                              "sha256:6432598541c0e3603d18cb1fd4e8946d5ad2772234a690e5bd242929c32289ee"
		"namespacelabs.dev/foundation/std/networking/gateway/server/configure":                        "sha256:1b21e06c1c3f506e4148152d3fe6ff48ab7fbdc7997e208f5b0b397ca03942a8"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/controller/img":                          "sha256:623154594a9edc30038330bef123354dc9c6ef6f74e6470609506d7f78a1f2cf"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/controller/tool":                         "sha256:93cf715829c24bebdf0876cac64a9d8b3112d13b236d68af88069d1003b32e3f"
		"namespacelabs.dev/foundation/std/runtime/kubernetes/kube-state-metrics/configure":            "sha256:d911ccfc7ee483c332ee7838af84aff411aa953b3565e03b23f16fca1df61163"
		"namespacelabs.dev/foundation/std/secrets/kubernetes":                                         "sha256:f88f89e60c8bb5059e9adbc5c1cd22d0780faa6be2a3b073283d30842f1387e6"
		"namespacelabs.dev/foundation/std/startup/testdriver":                                         "sha256:87ed023cec48cade0a5cdb4a433cc7859863cb0c7cddab6fd913249a1ead5ad5"
		"namespacelabs.dev/foundation/std/testdata/datastore/keygen":                                  "sha256:c6509339ab193bee825208494b5f79a1c58f3e190e03fe546dfc8b8cb796ee46"
		"namespacelabs.dev/foundation/std/web/http/configure":                                         "sha256:712eeb23a528f907df55f5edfd52e1683866db7271604eca9ec6b713d1af0821"
		"namespacelabs.dev/foundation/universe/aws/ecr/configure":                                     "sha256:165948f95ffa4765e67df2cd144a57d0334db3c6f13cb2d9f07ad9aaf16b584c"
		"namespacelabs.dev/foundation/universe/aws/irsa/prepare":                                      "sha256:0f66c98e52d900ccdf4d2b878976a1c2a447cc1cfbb7e9739b281f06e83c3129"
		"namespacelabs.dev/foundation/universe/aws/s3/internal/configure":                             "sha256:cd015460ae128f08e16f3d4273bb4d9438ed8f65a4b94afdff5d6042b6e55d75"
		"namespacelabs.dev/foundation/universe/aws/s3/internal/managebuckets/init":                    "sha256:a8e2852c25f94ba5bad5a10ba23e0286cd00944085fcd2ddb372d3a5caf43cad"
		"namespacelabs.dev/foundation/universe/db/maria/incluster/tool":                               "sha256:a1d50c07adaab5e486d3282c6f91fb8d4bbf97390878f50a269cfbc59fc80960"
		"namespacelabs.dev/foundation/universe/db/maria/internal/init":                                "sha256:83f10e9bbce60ffb41feb4358e305792afba363d70c5d70293bfed9c9751d2d6"
		"namespacelabs.dev/foundation/universe/db/maria/server/creds/tool":                            "sha256:1c4c1ec01a59a384ddca8297b663c84d62747338ab59c3c06d3e42b314003180"
		"namespacelabs.dev/foundation/universe/db/maria/server/img":                                   "sha256:0395a7a421ccfedd0fe83a213ec74a39c1d3c0b6185bd751c4f7b3472d6f8fb5"
		"namespacelabs.dev/foundation/universe/db/postgres/incluster/tool":                            "sha256:7d638650fc9a917cb205ebc956db46808beb3283209dbac1dbd1e7b31c6fc59b"
		"namespacelabs.dev/foundation/universe/db/postgres/internal/init":                             "sha256:5a2773eb3339154613b7f77b236bc20ad66d5e6b07c053845490ff8bbe4b48be"
		"namespacelabs.dev/foundation/universe/db/postgres/opaque/tool":                               "sha256:5e327b479e3c9ac63168d17b3ea7014477d946cc570141831511a595136c034b"
		"namespacelabs.dev/foundation/universe/db/postgres/server/creds/tool":                         "sha256:801b6543a1b8b2cc31811a93ec65b71f1672e991620b4cb6b679d6dcf809b68a"
		"namespacelabs.dev/foundation/universe/db/postgres/server/img":                                "sha256:f2d4f60aab2e2746bb9d0eb19b2dd40a05bc1aced783fd5e98a8f90f548ff04d"
		"namespacelabs.dev/foundation/universe/development/localstack/s3/internal/configure":          "sha256:3e22a7e178e1edc35a757a44b5f9b12e167cf7e5cc9f7609a1ae86d92b230cab"
		"namespacelabs.dev/foundation/universe/development/localstack/s3/internal/managebuckets/init": "sha256:a8f68f46d31cf5b791f13b145bb62ff359753ca7754f093d922dc12e1c1e991f"
		"namespacelabs.dev/foundation/universe/networking/tailscale/image":                            "sha256:046f157c620a3532510fb5d6000632a536462a078c42805e3aad461c1c59e951"
		"namespacelabs.dev/foundation/universe/storage/s3/internal/managebuckets":                     "sha256:5a4adc6caec3a281528a7a7b26cbbd49d4ab6c88fe85290670acb0a119e86b26"
		"namespacelabs.dev/foundation/universe/storage/s3/internal/prepare":                           "sha256:8449f07819d6fa75e9a218f443055778d9ce26a556c996b884961afd539156ff"
	}
	baseRepository: "us-docker.pkg.dev/foundation-344819/prebuilts/"
}
