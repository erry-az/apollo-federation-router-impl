install:
	curl -sSL https://rover.apollo.dev/nix/v0.9.1 | sh
gen:
	rover supergraph compose --config rover.config.yaml > supergraph-schema.graphql
dev:
	cargo run -- --hot-reload --config router.yaml --supergraph supergraph-schema.graphql --dev
