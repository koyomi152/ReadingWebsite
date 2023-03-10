package model

type Book struct {
	BookId      int    `json:"book_id"`
	Name        string `json:"name"`
	IsStar      bool   `json:"is_star"`
	Author      string `json:"author"`
	CommentNum  int    `json:"comment_num"`
	Score       int    `json:"score"`
	Cover       string `json:"cover"`
	PublishTime string `json:"publish_time"`
	Link        string `json:"link"`
}
