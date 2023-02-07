package service

import (
	"ReadingWebsite/dao"
	"ReadingWebsite/model"
)

func SearchComment(id int) (list []model.Comment, err error) {
	list, err = dao.SearchComment(id)
	return
}
func WriteComment(c model.Comment) (err error) {
	err = dao.WriteComment(c)
	return
}
func DeleteComment(id int) (err error) {
	err = dao.DeleteComment(id)
	return
}
func UpdateComment(id int, content string) (err error) {
	err = dao.UpdateComment(id, content)
	return
}
