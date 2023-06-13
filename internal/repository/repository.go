package repository

import (
	"github.com/KDias-code/internal/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(number, password string) (model.User, error)
	Update(userId int, input model.UpdateUserInput) error
	SmsCheck(verifyCode string, number string, user model.User) error
	RndSave(verifyCode string, number string, user model.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
