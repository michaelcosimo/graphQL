package structs

type Comment struct {
	ID      string `json:"id"`
	PostID  string `json:"postID"`
	Content string `json:"content"`
}
