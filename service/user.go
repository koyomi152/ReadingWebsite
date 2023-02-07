package service

import (
	"ReadingWebsite/dao"
	"ReadingWebsite/model"
	"math/rand"
)

func SearchUserByNickName(nickname string) (u model.User, err error) {
	u, err = dao.SearchUserByNickName(nickname)
	return
}
func SearchUserById(id int) (u model.User, err error) {
	u, err = dao.SearchUserById(id)
	return
}
func CreateUser(u model.User, password string) error {
	err := dao.CreateUser(u, password)
	return err
}
func SearchPasswordByNickName(nickname string) (password string, err error) {
	password, err = dao.SearchPasswordByNickName(nickname)
	return
}
func ModifyPassword(nickname string, NewPassword string) error {
	err := dao.ModifyPassword(nickname, NewPassword)
	return err
}
func CreateToken(nickname string, token string) error {
	err := dao.CreateToken(nickname, token)
	return err
}
func SearchTokenByNickname(nickname string) (token string, err error) {
	token, err = dao.SearchTokenByNickname(nickname)
	return token, err
}
func DeleteToken(nickname string) error {
	err := dao.DeleteToken(nickname)
	return err
}
func RandString(n int) (ret string) {
	allString := "qwertyuiopasdfghjklzxcvbnm0123456789"
	ret = ""
	for i := 0; i < n; i++ {
		r := rand.Intn(len(allString))
		ret = ret + allString[r:r+1]
	}
	return
}
func SearchToken(token string) (nickname string, err error) {
	nickname, err = dao.SearchToken(token)
	return nickname, err
}
