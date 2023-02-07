package api

import (
	"ReadingWebsite/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/register", Register)
	u := r.Group("/user")
	{
		u.POST("/token", Login)
		u.GET("token/refresh", middleware.AuthMiddleWare(), Refresh)
		u.PUT("/password", middleware.AuthMiddleWare(), Password)
		u.GET("/info/:user_id", middleware.AuthMiddleWare(), UserMessage)
	}
	m := r.Group("/book")
	{
		m.Use(middleware.AuthMiddleWare())
		m.GET("/list", BookList)
		m.GET("/search", BookSearch)
		m.PUT("/star", StarBook)
		m.GET("/label", BookLabel)
	}
	n := r.Group("/comment")
	{
		n.Use(middleware.AuthMiddleWare())
		m.GET("/:book_id", GetComment)
		m.POST("/:book_id", WriteComment)
		m.DELETE("/:comment_id", DeleteComment)
		m.PUT("/:comment_id", UpdateComment)
	}

	r.Run()
}
