package util

import (
	"ReadingWebsite/model"
	"github.com/gin-gonic/gin"
)

type respTemplate struct {
	Status int    `json:"Status"`
	Info   string `josn:"Info"`
}

var OK = respTemplate{
	Status: 10000,
	Info:   "success",
}

var ParamError = respTemplate{
	Status: 300,
	Info:   "params error",
}
var InternalError = respTemplate{
	Status: 500,
	Info:   "internal error",
}

func RespParamErr(c *gin.Context) {
	c.JSON(300, ParamError)
}
func RespInternalErr(c *gin.Context) {
	c.JSON(300, InternalError)
}
func RespRegister(c *gin.Context) {
	c.JSON(10000, OK)
}
func RespPassword(c *gin.Context) {
	c.JSON(10000, OK)
}
func RespInfo(c *gin.Context) {
	c.JSON(10000, OK)
}
func RespInfoUserId(c *gin.Context, u model.User) {
	c.JSON(10000, gin.H{
		"status": 10000,
		"info":   "success",
		"data": gin.H{
			"user": gin.H{
				"id":           u.Id,
				"avatar":       u.Avatar,
				"nickname":     u.Nickname,
				"introduction": u.Introduction,
				"phone":        u.Phone,
				"qq":           u.QQ,
				"gender":       u.Gender,
				"email":        u.Email,
				"birthday":     u.Birthday}}})
}
func RespToken(c *gin.Context, refreshToken string, token string) {
	c.JSON(10000, gin.H{
		"status": 10000,
		"info":   "success",
		"data": gin.H{
			"refresh_token": refreshToken,
			"token":         token}})
}

func RespUser(c *gin.Context, user model.User) {
	c.JSON(10000, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   user})
}
func NormErr(c *gin.Context, status int, info string) {
	c.JSON(300, gin.H{
		"status": status,
		"info":   info,
	})
}
