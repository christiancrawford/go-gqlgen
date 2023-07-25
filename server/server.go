package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/christiancrawford/go-gqlgen/graph/common"
	"github.com/christiancrawford/go-gqlgen/graph/generated"
	resolvers "github.com/christiancrawford/go-gqlgen/graph/resolvers"
)

const defaultPort = "4000"

// This code is starting up the GraphQL server:
func StartServer() {
	port := os.Getenv("PORT")

	// It gets the port to listen on from environment variable or uses a default.
	if port == "" {
		port = defaultPort
	}

	// It initializes the database connection.
	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	// It creates a GraphQL handler using gqlgen, passing in the generated schema and resolvers.
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	// It creates a context struct to hold the db connection and attaches it to each request.
	customCtx := &common.CustomContext{
		Database: db,
	}

	// It sets up routes for the GraphQL playground and the /query endpoint.
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", common.CreateContext(customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	// It starts the HTTP server listening on the configured port.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
