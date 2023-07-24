package handlers

import (
	"github.com/d3vtech/graphQL/database"
	"github.com/d3vtech/graphQL/structs"
	"github.com/graphql-go/graphql"
)

// Resolve comments for a specific post by postID
func ResolveComments(p graphql.ResolveParams) (interface{}, error) {
	postID := p.Args["postID"].(string)

	// Load comments for the specified postID
	comments := FetchCommentsFromDB(postID)

	return comments, nil
}

// Fetch comments from the data slice by postID
func FetchCommentsFromDB(postID string) []*structs.Comment {
	// Fetch comments from the data slice based on the provided postID
	var comments []*structs.Comment
	for _, comment := range database.CommentsData {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}
	return comments
}
