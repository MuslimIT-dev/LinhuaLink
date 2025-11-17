package repository

import (
	"LinhuaLink/backend/internal/model"
	"database/sql"
	"time"
)

type UserRepository interface {
	Signup(firstName, lastName, email, password string, birthDay time.Time) (model.User, error)
	Login(email string) (model.User, string, error)
	Me(userId int) (model.User, error)
}
type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Signup(firstName, lastName, email, password string, birthDay time.Time) (model.User, error) {
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

func (r *userRepo) Login(email string) (model.User, string, error) {
	var user model.User
	var password string
	user.Email = email
	err := r.db.QueryRow("SELECT id, first_name, last_name, password,birth_day, created_at FROM users WHERE email=$1", email).Scan(&user.ID, &user.FirstName, &user.LastName, &password, &user.BirthDay, &user.CreatedAt)
	if err != nil {
		return model.User{}, "", err
	}
	return user, password, nil
}

func (r *userRepo) Me(userId int) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT first_name,last_name,email,birth_day,created_at FROM users WHERE id=$1", userId).Scan(&user.FirstName, &user.LastName, &user.Email, &user.BirthDay, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
