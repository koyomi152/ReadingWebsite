package util

import (
	"ReadingWebsite/model"
	"github.com/gin-gonic/gin"
)

func RespGetComment(c *gin.Context, comment []model.Comment) {
	c.JSON(10000, gin.H{
		"info":    "success",
		"status":  10000,
		"comment": comment,
	})
}
func RespPostComment(c *gin.Context, id int) {
	c.JSON(10000, gin.H{
		"info":   "success",
		"status": 10000,
		"data":   id,
	})
}
func RespDeleteComment(c *gin.Context) {
	c.JSON(10000, gin.H{
		"info":   "success",
		"status": 10000,
	})
}
func RespUpdateComment(c *gin.Context) {
	c.JSON(10000, gin.H{
		"info":   "success",
		"status": 10000,
	})
}
