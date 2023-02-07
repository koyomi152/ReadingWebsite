package service

import (
	"ReadingWebsite/dao"
	"ReadingWebsite/model"
)

func BookList() (bookList []model.Book, err error) {
	bookList, err = dao.BookList()
	return
}

func BookSearch(BookName string) (b model.Book, err error) {
	b, err = dao.SearchBook(BookName)
	return
}
func StarBook(BookName string) error {
	err := dao.StarBook(BookName)
	return err
}
func BookLabel(label string) (BookList []model.Book, err error) {
	BookList, err = dao.BookLabel(label)
	return
}
