package api

import (
	"ReadingWebsite/model"
	"ReadingWebsite/service"
	"ReadingWebsite/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func GetComment(c *gin.Context) {
	bookId := c.Query("book_id")
	Id, _ := strconv.Atoi(bookId)
	list, err := service.SearchComment(Id)
	if err != nil {
		log.Println(err)
	}
	util.RespGetComment(c, list)
}
func WriteComment(c *gin.Context) {
	token := c.PostForm("token")
	nickname, _ := service.SearchToken(token)
	user, _ := service.SearchUserByNickName(nickname)
	var Comment model.Comment
	Comment.UserId = user.Id
	Comment.Avatar = user.Avatar
	Comment.Nickname = user.Nickname
	Comment.PublishTime = time.Now().Unix()
	bookId := c.Query("book_id")
	Id, _ := strconv.Atoi(bookId)
	Comment.BookId = Id
	Comment.PraiseCount = 0
	Comment.IsPraised = false
	Comment.IsFocus = false
	content := c.PostForm("content")
	Comment.Content = content
	err := service.WriteComment(Comment)
	if err != nil {
		log.Println(err)
	}
	util.RespPostComment(c, Comment.PostId)
}
func DeleteComment(c *gin.Context) {
	commentId := c.Query("comment_id")
	Id, _ := strconv.Atoi(commentId)
	err := service.DeleteComment(Id)
	if err != nil {
		log.Println(err)
	}
	util.RespDeleteComment(c)
}
func UpdateComment(c *gin.Context) {
	commentId := c.Query("comment_id")
	Id, _ := strconv.Atoi(commentId)
	content := c.PostForm("content")
	err := service.UpdateComment(Id, content)
	if err != nil {
		log.Println(err)
	}
	util.RespUpdateComment(c)
}
