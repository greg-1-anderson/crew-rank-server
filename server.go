package main

import (
	"log"
	"net/http"
	"os"
	"twta/crew-rank-server/graph"
	"twta/crew-rank-server/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8443"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)
	certFile := "certs/crewrank.crt"
	keyFile := "certs/crewrank.key"
	log.Fatal(http.ListenAndServeTLS(":"+port, certFile, keyFile, nil))
}
