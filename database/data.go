package database

import (
	"github.com/d3vtech/graphQL/structs"
)

var PostsData = []*structs.Post{
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

var CommentsData = []*structs.Comment{
	{ID: "1", PostID: "1", Content: "Comment 1 on Post 1"},
	{ID: "2", PostID: "1", Content: "Comment 2 on Post 1"},
	{ID: "3", PostID: "2", Content: "Comment 1 on Post 2"},
	{ID: "4", PostID: "3", Content: "Comment 1 on Post 3"},
	{ID: "5", PostID: "3", Content: "Comment 2 on Post 3"},
	{ID: "6", PostID: "4", Content: "Comment 1 on Post 4"},
	// Add more comments here
}
