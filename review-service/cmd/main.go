package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/erry-az/go-gql-federation/review-service/data"
	"github.com/erry-az/go-gql-federation/review-service/graph"
	"github.com/erry-az/go-gql-federation/review-service/graph/generated"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8088"

func main() {
	port := os.Getenv("REVIEW_PORT")
	if port == "" {
		port = defaultPort
	}

	data.InitReviewers()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
