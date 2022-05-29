package model

type User struct {
	Id       int16
	UserName string `json:"username"`
	Password string `json:"password"`
}
