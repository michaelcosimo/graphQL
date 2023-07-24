package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/d3vtech/graphQL/graph"
	"github.com/d3vtech/graphQL/middleware"
)

func main() {
	// Define the GraphQL schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graph.RootQuery,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Create a new GraphQL handler with the schema
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// // Wrap in authorization
	a := middleware.AuthMiddleware(h)

	// // Wrap the GraphQL handler with the caching middleware
	// c := middleware.CachingMiddleware(a)

	// // Wrap the GraphQL handler with the stitching middleware (placeholder for potential future use)
	// s := middleware.StitchingMiddleware(c)

	// Wrap the GraphQL handler with the authentication and authorization middleware
	http.Handle("/graphql", a)

	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
