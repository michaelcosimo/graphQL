package structs

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

// Post represents a post in the application
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
