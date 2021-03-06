package main

import (
	"log"
	"net/http"
	"os"
	"twta/crew-rank-server/graph"
	"twta/crew-rank-server/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8443"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Public API: allow all callers
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				// return r.Host == "example.org"
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)
	certFile := "certs/crewrank.crt"
	keyFile := "certs/crewrank.key"
	log.Fatal(http.ListenAndServeTLS(":"+port, certFile, keyFile, router))
}
