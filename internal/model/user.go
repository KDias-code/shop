package model

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Number   string `json:"number" binding:"required"`
	Password string `json:"password" binding:"required"`
}
