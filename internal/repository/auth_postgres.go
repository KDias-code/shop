package repository

import (
	"errors"
	"fmt"
	"github.com/KDias-code/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, number, password, type) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Number, user.Password, user.Type)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(number, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE number=$1 AND password=$2", userTable)
	err := r.db.Get(&user, query, number, password)

	return user, err
}

func (r *AuthPostgres) Update(userId int, input model.UpdateUserInput) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", userTable)
	_, err := r.db.Exec(query, *input.Name, userId)
	return err
}

func (r *AuthPostgres) RndSave(verifyCode int, number string, user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (sms) VALUES ($1) WHERE number =$2 RETURNING id", userTable)
	row := r.db.QueryRow(query, verifyCode, number)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) SmsCheck(verifyCode int, number string, user model.User) error {
	var storedCode int
	query := fmt.Sprintf("SELECT sms FROM %s WHERE number = $1", userTable)
	err := r.db.Get(&storedCode, query, number)
	if err != nil {
		return err
	}

	if storedCode != verifyCode {
		return errors.New("verification code mismatch")
	}

	return nil
}
