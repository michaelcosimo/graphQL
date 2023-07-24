package handlers

import (
	"context"
	"sort"

	"github.com/d3vtech/graphQL/database"
	"github.com/d3vtech/graphQL/structs"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql"
)

// Helper function to get the edges for the posts
func GetPostEdges(posts []*structs.Post) []*structs.PostEdge {
	edges := make([]*structs.PostEdge, len(posts))
	for i, post := range posts {
		edges[i] = &structs.PostEdge{
			Cursor: GetPostCursor(post),
			Node:   post,
		}
	}
	return edges
}

// Helper function to get the cursor for a post
func GetPostCursor(post *structs.Post) string {
	return post.ID
}

func ResolvePostIDs(p graphql.ResolveParams) (interface{}, error) {
	ids := p.Args["ids"].([]string)

	// Load posts from the data slice by IDs
	posts := fetchPostsFromDB(ids)

	return posts, nil
}

// Fetch posts from the data slice by IDs
func fetchPostsFromDB(ids []string) []*structs.Post {
	// Fetch posts from the data slice based on the provided IDs
	var posts []*structs.Post
	for _, id := range ids {
		for _, post := range database.PostsData {
			if post.ID == id {
				posts = append(posts, post)
				break
			}
		}
	}
	return posts
}

// FetchPosts fetches posts from the data slice by their IDs
func FetchPosts(ctx context.Context, keys dataloader.Keys) []*structs.Post {
	ids := make([]string, len(keys))
	for i, key := range keys {
		ids[i] = key.String()
	}

	// Fetch posts from the data slice by their IDs
	var posts []*structs.Post
	for _, id := range ids {
		// Find the post with the given ID in the PostsData slice
		for _, post := range database.PostsData {
			if post.ID == id {
				posts = append(posts, post)
				break
			}
		}
	}

	// Sort the fetched posts to match the order of keys
	sort.Slice(posts, func(i, j int) bool {
		return indexOf(ids, posts[i].ID) < indexOf(ids, posts[j].ID)
	})

	return posts
}

// Helper function to find the index of a string in a slice
func indexOf(slice []string, item string) int {
	for i, s := range slice {
		if s == item {
			return i
		}
	}
	return -1
}

func ResolvePost(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)

	// Fetch the post data from the data slice
	var post *structs.Post
	for _, p := range database.PostsData {
		if p.ID == id {
			post = p
			break
		}
	}

	return post, nil
}
