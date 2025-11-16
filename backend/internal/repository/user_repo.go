package repository

import (
	"LinhuaLink/backend/internal/model"
	"database/sql"
	"time"
)

type UserRepository interface {
	Signup(firstName, lastName, email, password string, birthDay time.Time) (model.User, error)
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
	user.Password = password
	user.BirthDay = birthDay
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
