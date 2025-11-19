package repository

import (
	"LinhuaLink/backend/internal/model"
	"database/sql"
	"time"
)

type AuthRepository interface {
	Signup(firstName, lastName, email, password string, birthDay time.Time) (model.User, error)
	Login(email string) (model.User, string, error)
}
type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Signup(firstName, lastName, email, password string, birthDay time.Time) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(
		`INSERT INTO users (first_name, last_name, email, password, birth_day) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id,created_at`,
		firstName, lastName, email, password, birthDay,
	).Scan(&user.ID, &user.CreatedAt)

	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	user.BirthDay = birthDay
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *authRepository) Login(email string) (model.User, string, error) {
	var user model.User
	var password string
	user.Email = email
	err := r.db.QueryRow("SELECT id, first_name, last_name, password,birth_day, created_at FROM users WHERE email=$1", email).Scan(&user.ID, &user.FirstName, &user.LastName, &password, &user.BirthDay, &user.CreatedAt)
	if err != nil {
		return model.User{}, "", err
	}
	return user, password, nil
}
