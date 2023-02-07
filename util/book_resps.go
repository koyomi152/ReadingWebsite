package util

import (
	"ReadingWebsite/model"
	"github.com/gin-gonic/gin"
)

func RespBookList(c *gin.Context, bookList []model.Book) {
	c.JSON(10000, gin.H{
		"status": 10000,
		"info":   "success",
		"date": gin.H{
			"books": bookList}})
}
func RespBook(c *gin.Context, b model.Book) {
	c.JSON(10000, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   b,
	})
}
func RespStar(c *gin.Context) {
	c.JSON(10000, gin.H{
		"status": 10000,
		"info":   "success",
	})
}
func RespBookByLabel(c *gin.Context, BookList []model.Book) {
	c.JSON(10000, gin.H{
		"status": 10000,
		"info":   "success",
		"data": gin.H{
			"books": BookList,
		}})
}
