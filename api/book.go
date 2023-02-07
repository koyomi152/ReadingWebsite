package api

import (
	"ReadingWebsite/service"
	"ReadingWebsite/util"
	"github.com/gin-gonic/gin"
	"log"
)

func BookList(c *gin.Context) {
	bookList, err := service.BookList()
	if err != nil {
		log.Println(err)
	}
	util.RespBookList(c, bookList)
}
func BookSearch(c *gin.Context) {
	BookName := c.PostForm("BookName")
	b, err := service.BookSearch(BookName)
	if err != nil {
		log.Println(err)
	}
	util.RespBook(c, b)
}
func StarBook(c *gin.Context) {
	BookName := c.PostForm("BookName")
	err := service.StarBook(BookName)
	if err != nil {
		log.Println(err)
	}
	util.RespStar(c)
}
func BookLabel(c *gin.Context) {
	label := c.PostForm("label")
	BookList, err := service.BookLabel(label)
	if err != nil {
		log.Println(err)
	}
	util.RespBookByLabel(c, BookList)
}
