package model

type User struct {
	Id           int    `json:"Id"`
	Avatar       string `json:"Avatar"`
	Nickname     string `json:"Nickname"`
	Introduction string `json:"Introduction"`
	Phone        int    `json:"Phone"`
	QQ           int    `json:"QQ"`
	Gender       string `json:"Gender"`
	Email        string `json:"Email"`
	Birthday     string `json:"Birthday"`
}
