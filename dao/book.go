package dao

import (
	"ReadingWebsite/model"
	"log"
)

func BookList() (bookList []model.Book, err error) {
	row, err := DB.Query("select book_id,name,is_star,author,comment_num,score,cover,publish_time,link from book")
	if err = row.Err(); row.Err() != nil {
		return
	}
	bookList = []model.Book{}
	for row.Next() {
		var b model.Book
		err = row.Scan(&b.BookId, &b.Name, &b.IsStar, &b.Author, &b.CommentNum, &b.Score, &b.Cover, &b.PublishTime, &b.Link)
		if err != nil {
			log.Println(err)
			return
		}
		bookList = append(bookList, b)
	}
	return
}
func SearchBook(BookName string) (b model.Book, err error) {
	row := DB.QueryRow("select  book_id,name,is_star,author,comment_num,score,cover,publish_time,link from book where name=?", BookName)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&b.BookId, &b.Name, &b.IsStar, &b.Author, &b.CommentNum, &b.Score, &b.Cover, &b.PublishTime, &b.Link)
	if err != nil {
		log.Println(err)
	}
	return
}
func StarBook(BookName string) (err error) {
	_, err = DB.Exec("update book set is_star=? where name=?", 1, BookName)
	return
}
func BookLabel(label string) (BookList []model.Book, err error) {
	row, err := DB.Query("select book_id,name,is_star,author,comment_num,score,cover,publish_time,link from book where label=?", label)
	if err = row.Err(); err != nil {
		return
	}
	BookList = []model.Book{}
	for row.Next() {
		var b model.Book
		err = row.Scan(&b.BookId, &b.Name, &b.IsStar, &b.Author, &b.CommentNum, &b.Score, &b.Cover, &b.PublishTime, &b.Link)
		if err != nil {
			log.Println(err)
		}
		BookList = append(BookList, b)
	}
	return
}
