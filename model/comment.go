package model

type Comment struct {
	PostId      int    `json:"post_id"`
	PublishTime int64  `json:"publish_time"`
	Content     string `json:"content"`
	UserId      int    `json:"user_id"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	PraiseCount int    `json:"praise_count"`
	IsPraised   bool   `json:"is_praised"`
	IsFocus     bool   `json:"is_focus"`
	BookId      int    `json:"book_id"`
}
