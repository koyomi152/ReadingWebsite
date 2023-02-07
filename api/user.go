package api

import (
	"ReadingWebsite/model"
	"ReadingWebsite/service"
	"ReadingWebsite/util"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Register(c *gin.Context) {
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	gender := c.PostForm("gender")
	qq := c.PostForm("qq")
	email := c.PostForm("email")
	avatar := c.PostForm("avatar")
	introduction := c.PostForm("introduction")
	phone := c.PostForm("phone")
	birthday := c.PostForm("birthday")
	QQ, _ := strconv.Atoi(qq)
	Phone, _ := strconv.Atoi(phone)
	if len(nickname) == 0 || len(nickname) > 20 {

		return
	}
	if len(password) < 6 {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByNickName(nickname)
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalErr(c)
		return
	}
	if u.Nickname != "" {
		util.RespParamErr(c)
		return
	}
	err = service.CreateUser(model.User{
		Nickname:     nickname,
		Gender:       gender,
		QQ:           QQ,
		Email:        email,
		Avatar:       avatar,
		Introduction: introduction,
		Phone:        Phone,
		Birthday:     birthday,
	}, password)
	if err != nil {
		util.RespParamErr(c)
		return
	}
	util.RespRegister(c)
}
func Login(c *gin.Context) {
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	if len(nickname) == 0 || len(nickname) > 20 {

		return
	}
	if len(password) < 6 {
		util.RespParamErr(c)
		return
	}
	_, err := service.SearchUserByNickName(nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
			return
		} else {
			log.Printf("search user error:%v", err)
		}
		return
	}
	Password, err := service.SearchPasswordByNickName(nickname)
	if Password != password {
		util.NormErr(c, 300, "密码错误")
	}
	token, err := service.SearchTokenByNickname(nickname)
	if err == nil {
		service.DeleteToken(nickname)
	}
	NewToken := service.RandString(6)
	err = service.CreateToken(nickname, NewToken)
	if err != nil {
		util.RespParamErr(c)
	}
	util.RespToken(c, token, NewToken)
}
func Refresh(c *gin.Context) {
	token := c.PostForm("token")
	nickname, err := service.SearchToken(token)
	token, err = service.SearchTokenByNickname(nickname)
	if err == nil {
		service.DeleteToken(nickname)
	}
	NewToken := service.RandString(6)
	err = service.CreateToken(nickname, NewToken)
	if err != nil {
		util.RespParamErr(c)
	}
	util.RespToken(c, token, NewToken)
}
func Password(c *gin.Context) {
	token := c.PostForm("token")
	password := c.PostForm("password")
	newPassword := c.PostForm("NewPassword")
	nickname, _ := service.SearchToken(token)
	num, _ := service.SearchPasswordByNickName(nickname)
	if num != password {
		util.NormErr(c, 300, "密码错误")
	}
	err := service.ModifyPassword(nickname, newPassword)
	if err != nil {
		println(err)
	}
	util.RespPassword(c)
}
func UserMessage(c *gin.Context) {
	id := c.Query("user_id")
	Id, _ := strconv.Atoi(id)
	u, err := service.SearchUserById(Id)
	if err != nil {
		println(err)
	}
	util.RespInfoUserId(c, u)
}
