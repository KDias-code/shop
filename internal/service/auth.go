package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/KDias-code/internal/model"
	"github.com/KDias-code/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt     = "asdardirqweo231koads"
	signKey  = "fjdisi321ewqieqw"
	tokenTTl = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(number, password string) (string, error) {
	user, err := s.repo.GetUser(number, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid singing method")
		}

		return []byte(signKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) Update(userId int, input model.UpdateUserInput) error {
	return s.repo.Update(userId, input)
}

func (s *AuthService) RndSave(verifyCode int, number string, user model.User) (int, error) {
	return s.repo.RndSave(verifyCode, number, user)
}

func (s *AuthService) SmsCheck(verifyCode string, number string, user model.User) error {
	return s.repo.SmsCheck(verifyCode, number, user)
}
