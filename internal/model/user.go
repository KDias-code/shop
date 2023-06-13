package model

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Number   string `db:"number"`
	Surname  string `db:"surname"`
	Password string `db:"password"`
	Type     string `db:"type"`
	SMS      string `db:"sms"`
}
