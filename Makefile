user/gen:
	cd user-service && go run github.com/99designs/gqlgen
review/gen:
	cd review-service && go run github.com/99designs/gqlgen
router/gen:
	rover supergraph compose --config graphql-router/rover.config.yaml > graphql-router/supergraph-schema.graphql

user/dev:
	cd user-service && go run cmd/main.go &
review/dev:
	cd review-service && go run cmd/main.go &
router/dev:
	make router/gen && cd graphql-router && cargo run -- --hot-reload --config router.yaml --supergraph supergraph-schema.graphql --dev --log debug
dev:
	make user/dev
	make review/dev
stop:
	kill $$(ps -ef | grep tmp/go-build | grep -v grep | awk '{print $$2}')