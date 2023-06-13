package service

import (
	"github.com/KDias-code/internal/model"
	"github.com/KDias-code/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(number, password string) (string, error)
	ParseToken(token string) (int, error)
	Update(userId int, input model.UpdateUserInput) error
	SmsCheck(verifyCode string, number string, user model.User) error
	RndSave(verifyCode int, number string, user model.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
