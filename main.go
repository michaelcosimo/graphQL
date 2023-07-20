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
					// Retrieve the first and after cursor values from the arguments
					first := p.Args["first"].(int)
					afterCursor := p.Args["after"].(string)

					// Implement your cursor-based pagination logic here
					// Find the index of the item with the provided after cursor
					startIndex := 0
					for i, post := range postsData {
						if post.ID == afterCursor {
							startIndex = i + 1
							break
						}
					}

					// Return the paginated results based on the first and after cursor
					endIndex := startIndex + first
					if endIndex > len(postsData) {
						endIndex = len(postsData)
					}

					return struct {
						Edges    []*PostEdge `json:"edges"`
						PageInfo PageInfo    `json:"pageInfo"`
					}{
						Edges: getPostEdges(postsData[startIndex:endIndex]),
						PageInfo: PageInfo{
							HasNextPage: endIndex < len(postsData),
							EndCursor:   getPostCursor(postsData[endIndex-1]),
						},
					}, nil
				},
			},
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
			},
			"pageInfo": &graphql.Field{
				Type: pageInfo,
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
			},
			"node": &graphql.Field{
				Type: postType,
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
			},
			"endCursor": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Helper function to get the edges for the posts
func getPostEdges(posts []*Post) []*PostEdge {
	edges := make([]*PostEdge, len(posts))
	for i, post := range posts {
		edges[i] = &PostEdge{
			Cursor: getPostCursor(post),
			Node:   post,
		}
	}
	return edges
}

// Helper function to get the cursor for a post
func getPostCursor(post *Post) string {
	return post.ID
}

// PostEdge represents an edge in the pagination
type PostEdge struct {
	Cursor string `json:"cursor"`
	Node   *Post  `json:"node"`
}

// PageInfo represents page information
type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

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
