package dao

import (
	"ReadingWebsite/model"
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ReadingWebsite?charset=uft8mb4&loc=Local&parseTime=true")
	if err != nil {
		log.Fatalf("connect mysql error:%v", err)
		return
	}
	DB = db
}
func SearchUserByNickName(nickname string) (u model.User, err error) {
	row := DB.QueryRow("select id,gender,nickname,qq,birthday,email,avatar,introduction,phone from user where nickname=?", nickname)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.Gender, &u.Nickname, &u.QQ, &u.Birthday, &u.Email, &u.Avatar, &u.Introduction, &u.Phone)
	return
}
func SearchUserById(id int) (u model.User, err error) {
	row := DB.QueryRow("select id,gender,nickname,qq,birthday,email,avatar,introduction,phone from user where id=?", id)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.Gender, &u.Nickname, &u.QQ, &u.Birthday, &u.Email, &u.Avatar, &u.Introduction, &u.Phone)
	return
}
func CreateUser(u model.User, password string) (err error) {
	_, err = DB.Exec("insert into user(gender,nickname,qq,birthday,email,avatar,introduction,phone,password)values (?,?,?,?,?,?,?,?,?)", u.Gender, u.Nickname, u.QQ, u.Birthday, u.Email, u.Avatar, u.Introduction, u.Phone, password)
	return
}
func SearchPasswordByNickName(nickname string) (password string, err error) {
	row := DB.QueryRow("select password from user where nickname=?", nickname)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(password)
	return
}
func ModifyPassword(nickname string, NewPassword string) (err error) {
	_, err = DB.Exec("update user set password=? where nickname=?", NewPassword, nickname)
	return
}

func CreateToken(nickname string, token string) (err error) {
	_, err = DB.Exec("insert into Token(nickname,token)values (?,?)", nickname, token)
	return
}
func SearchTokenByNickname(nickname string) (token string, err error) {
	row := DB.QueryRow("select token from Token where nickname=?", nickname)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(token)
	return
}
func DeleteToken(nickname string) (err error) {
	_, err = DB.Exec("delete from Token where nickname=?", nickname)
	return
}
func SearchToken(token string) (nickname string, err error) {
	row := DB.QueryRow("select nickname from Token where token=?", token)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(nickname)
	return nickname, err
}
