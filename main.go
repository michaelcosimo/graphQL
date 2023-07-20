package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Post represents a post in the application
type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// postsData represents a dummy data store for posts
var postsData = []*Post{
	{ID: "1", Title: "Post 1", Content: "Content of Post 1"},
	{ID: "2", Title: "Post 2", Content: "Content of Post 2"},
	{ID: "3", Title: "Post 3", Content: "Content of Post 3"},
	// Add more posts here
}

// Define the GraphQL type for Post
var postType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":      &graphql.Field{Type: graphql.ID},
			"title":   &graphql.Field{Type: graphql.String},
			"content": &graphql.Field{Type: graphql.String},
		},
	},
)

// Define a connection type for pagination
var postConnection = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostConnection",
		Fields: graphql.Fields{
			"edges": &graphql.Field{
				Type: graphql.NewList(postEdge),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Dummy resolve function to return the list of post edges
					// based on the provided pagination arguments (not implemented in this example)
					return postsData, nil
				},
			},
			"pageInfo": &graphql.Field{
				Type: pageInfo,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Dummy resolve function to return the page info
					// based on the provided pagination arguments (not implemented in this example)
					return map[string]interface{}{
						"hasNextPage": false,
						"endCursor":   nil,
					}, nil
				},
			},
		},
	},
)

// Define an edge type for pagination
var postEdge = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostEdge",
		Fields: graphql.Fields{
			"cursor": &graphql.Field{
				Type: graphql.ID,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Dummy resolve function to return the cursor for the current post (not implemented in this example)
					// You can implement actual logic to retrieve the cursor value for each post
					return "cursor_value", nil
				},
			},
			"node": &graphql.Field{
				Type: postType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Resolve and return the current post based on the provided cursor (pagination) (implemented in this example)
					post := p.Source.(*Post)
					return post, nil
				},
			},
		},
	},
)

// Define the page info type for pagination
var pageInfo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PageInfo",
		Fields: graphql.Fields{
			"hasNextPage": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Dummy resolve function to return the hasNextPage value (not implemented in this example)
					return false, nil
				},
			},
			"endCursor": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Dummy resolve function to return the endCursor value (not implemented in this example)
					return nil, nil
				},
			},
		},
	},
)

func main() {
	// Define the root query
	// ...

	// Define the root query
	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"posts": &graphql.Field{
					Type: postConnection,
					Args: graphql.FieldConfigArgument{
						"first": &graphql.ArgumentConfig{
							Type:         graphql.Int,
							DefaultValue: 10,
						},
						"after": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// Implement the resolve function to return the posts with pagination
						// based on the provided "first" and "after" arguments
						first := p.Args["first"].(int)
						afterCursor := p.Args["after"].(string)

						// Implement your pagination logic here
						// For simplicity, we'll return a subset of posts from the dummy data store
						var startIndex int
						for i, post := range postsData {
							if post.ID == afterCursor {
								startIndex = i + 1
								break
							}
						}

						if startIndex >= len(postsData) {
							return nil, nil
						}

						endIndex := startIndex + first
						if endIndex > len(postsData) {
							endIndex = len(postsData)
						}

						return postsData[startIndex:endIndex], nil
					},
				},
			},
		},
	)

	// ...

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
