package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Post represents a post in the application
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// postsData represents a dummy data store for posts
var postsData = []*Post{
	{ID: "1", Title: "Post 1"},
	{ID: "2", Title: "Post 2"},
	{ID: "3", Title: "Post 3"},
	{ID: "4", Title: "Post 4"},
	{ID: "5", Title: "Post 5"},
	{ID: "6", Title: "Post 6"},
	{ID: "7", Title: "Post 7"},
	{ID: "8", Title: "Post 8"},
	{ID: "9", Title: "Post 9"},
	{ID: "10", Title: "Post 10"},
	// Add more posts here
}

// Define the GraphQL type for Post
var postType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.ID},
			"title": &graphql.Field{Type: graphql.String},
		},
	},
)

// Define the root query
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"posts": &graphql.Field{
				Type: graphql.NewList(postType),
				Args: graphql.FieldConfigArgument{
					"offset": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
					},
					"limit": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 10,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Retrieve the offset and limit values from the arguments
					offset := p.Args["offset"].(int)
					limit := p.Args["limit"].(int)

					// Implement your pagination logic here
					startIndex := offset
					endIndex := offset + limit
					if endIndex > len(postsData) {
						endIndex = len(postsData)
					}

					return postsData[startIndex:endIndex], nil
				},
			},
		},
	},
)

func main() {
	// Define the GraphQL schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
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

	// Serve the GraphQL endpoint
	http.Handle("/graphql", h)
	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
