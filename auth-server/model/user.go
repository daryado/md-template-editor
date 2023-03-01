package model

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"login" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}
