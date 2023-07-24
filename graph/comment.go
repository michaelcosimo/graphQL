package graph

import (
	"github.com/d3vtech/graphQL/handlers"
	"github.com/graphql-go/graphql"
)

var CommentSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: commentType,
	},
)

// Define the root query fields for comments
var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Args: graphql.FieldConfigArgument{
					"postID": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: handlers.ResolveComments,
			},
		},
	},
)

// Define the GraphQL type for Comment
var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id":      &graphql.Field{Type: graphql.String},
			"postID":  &graphql.Field{Type: graphql.String},
			"content": &graphql.Field{Type: graphql.String},
		},
	},
)
