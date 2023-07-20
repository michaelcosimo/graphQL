package main

import (
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Book represents a book in the library
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	IsAvailable bool      `json:"isAvailable"`
	RentedBy    string    `json:"rentedBy,omitempty"`
	RentedAt    time.Time `json:"rentedAt,omitempty"`
}

// Define GraphQL types for Book
var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.String},
			"title":       &graphql.Field{Type: graphql.String},
			"author":      &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"price":       &graphql.Field{Type: graphql.Float},
			"isAvailable": &graphql.Field{Type: graphql.Boolean},
			"rentedBy":    &graphql.Field{Type: graphql.String},
			"rentedAt":    &graphql.Field{Type: graphql.DateTime},
		},
	},
)

// Define the root query
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Retrieve the book by ID
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, nil
					}

					// Retrieve book data from the database or other sources
					// Here, we're returning dummy data for demonstration purposes
					return &Book{
						ID:          id,
						Title:       "Book Title",
						Author:      "Book Author",
						Description: "Book Description",
						Price:       9.99,
						IsAvailable: true,
						RentedBy:    "John Doe",
						RentedAt:    time.Now(),
					}, nil
				},
			},
		},
	},
)

// Define the root mutation
var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"rentBook": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"rentedBy": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Retrieve the book by ID
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, nil
					}

					// Retrieve additional input
					rentedBy, _ := p.Args["rentedBy"].(string)

					// Rent the book
					book := &Book{
						ID:          id,
						Title:       "Book Title",
						Author:      "Book Author",
						Description: "Book Description",
						Price:       9.99,
						IsAvailable: false,
						RentedBy:    rentedBy,
						RentedAt:    time.Now(),
					}

					// Persist the rented book data in the database or other sources
					// Here, we're simply returning the rented book for demonstration purposes
					return book, nil
				},
			},
		},
	},
)

func main() {
	// Define the GraphQL schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery,
			Mutation: rootMutation,
		},
	)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Create a new GraphQL handler with the schema
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Serve the GraphQL endpoint
	http.Handle("/graphql", h)
	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
