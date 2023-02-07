package middleware

import (
	"ReadingWebsite/service"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("Token")
		_, err := service.SearchToken(token)
		if err != nil {
			log.Printf("search user error:%v", err)
		}
		c.Next()
		return
	}
}
