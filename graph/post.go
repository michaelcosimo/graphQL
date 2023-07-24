package graph

import (
	"github.com/d3vtech/graphQL/handlers"
	"github.com/d3vtech/graphQL/structs"
	"github.com/graphql-go/graphql"
)

// Define the GraphQL schemas for posts and comments
var PostSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: PostType,
	},
)

// Define the GraphQL type for Post
var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.ID},
			"title": &graphql.Field{Type: graphql.String},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Fetch comments for the current post
					post, _ := p.Source.(*structs.Post)
					return handlers.FetchCommentsFromDB(post.ID), nil
				},
			},
		},
	},
)

// Define an edge type for pagination
var PostEdge = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostEdge",
		Fields: graphql.Fields{
			"cursor": &graphql.Field{
				Type: graphql.ID,
			},
			"node": &graphql.Field{
				Type: PostType,
			},
		},
	},
)

// Define the page info type for pagination
var PageInfo = graphql.NewObject(
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

// Define a connection type for pagination
var PostConnection = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostConnection",
		Fields: graphql.Fields{
			"edges": &graphql.Field{
				Type: graphql.NewList(PostEdge),
			},
			"pageInfo": &graphql.Field{
				Type: PageInfo,
			},
		},
	},
)
