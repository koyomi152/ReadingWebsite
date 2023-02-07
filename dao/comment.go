package dao

import (
	"ReadingWebsite/model"
	"log"
)

func SearchComment(id int) (list []model.Comment, err error) {
	row, err := DB.Query("select post_id,publish_time,content,user_id, avatar,nickname,praise_count,is_praised,is_focus from comment where book_id=?", id)
	if err = row.Err(); err != nil {
		return
	}
	list = []model.Comment{}
	for row.Next() {
		var c model.Comment
		err := row.Scan(&c.PostId, &c.PublishTime, &c.Content, &c.UserId, &c.Avatar, &c.Nickname, &c.PraiseCount, &c.IsPraised, &c.IsFocus)
		if err != nil {
			log.Println(err)
		}
	}
	return
}
func WriteComment(c model.Comment) (err error) {
	_, err = DB.Exec("insert into comment(post_id,book_id,publish_time,content,user_id,avatar,nickname,praise_count,is_praised,is_focus) values(?,?,?,?,?,?,?,?,?,?)", c.PostId, c.BookId, c.PublishTime, c.Content, c.UserId, c.Avatar, c.Nickname, c.PraiseCount, c.IsPraised, c.IsFocus)
	return
}
func DeleteComment(id int) (err error) {
	_, err = DB.Exec("delete from comment where psot_id=?", id)
	return
}
func UpdateComment(id int, content string) (err error) {
	_, err = DB.Exec("update comment set content=? where post_id=?", content, id)
	return
}
