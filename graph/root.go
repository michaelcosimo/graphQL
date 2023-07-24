package graph

import (
	"errors"

	"github.com/d3vtech/graphQL/confiq"
	"github.com/d3vtech/graphQL/database"
	"github.com/d3vtech/graphQL/handlers"
	"github.com/d3vtech/graphQL/structs"
	"github.com/graphql-go/graphql"
)

// Define the root query
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"posts": &graphql.Field{
				Type: PostConnection,
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
					// Retrieve the user role from the request context
					role, ok := p.Context.Value("role").(string)
					if !ok {
						return nil, errors.New("unauthorized")
					}

					// Implement your authorization logic here to restrict access to certain queries
					// For example, allow only admins to access all posts
					if role == confiq.RoleClient {
						return nil, errors.New("unauthorized: access denied")
					}

					// Retrieve the first and after cursor values from the arguments
					first := p.Args["first"].(int)
					afterCursor := p.Args["after"].(string)

					// Implement your cursor-based pagination logic here
					// Find the index of the item with the provided after cursor
					startIndex := 0
					for i, post := range database.PostsData {
						if post.ID == afterCursor {
							startIndex = i + 1
							break
						}
					}

					// Return the paginated results based on the first and after cursor
					endIndex := startIndex + first
					if endIndex > len(database.PostsData) {
						endIndex = len(database.PostsData)
					}

					return struct {
						Edges    []*structs.PostEdge `json:"edges"`
						PageInfo structs.PageInfo    `json:"pageInfo"`
					}{
						Edges: handlers.GetPostEdges(database.PostsData[startIndex:endIndex]),
						PageInfo: structs.PageInfo{
							HasNextPage: endIndex < len(database.PostsData),
							EndCursor:   handlers.GetPostCursor(database.PostsData[endIndex-1]),
						},
					}, nil
				},
			},
			"post": &graphql.Field{
				Type: PostType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: handlers.ResolvePost,
			},
			"post_ids": &graphql.Field{
				Type: graphql.NewList(PostType),
				Args: graphql.FieldConfigArgument{
					"ids": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
					},
				},
				Resolve: handlers.ResolvePostIDs,
			},
			"posts_offset": &graphql.Field{
				Type: graphql.NewList(PostType),
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
					if endIndex > len(database.PostsData) {
						endIndex = len(database.PostsData)
					}

					return database.PostsData[startIndex:endIndex], nil
				},
			},
		},
	},
)
