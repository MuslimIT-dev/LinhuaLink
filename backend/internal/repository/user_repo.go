package repository

import (
	"LinhuaLink/backend/internal/model"
	"database/sql"
)

type UserRepository interface {
	Me(userId int) (model.User, error)
}
type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Me(userId int) (model.User, error) {
	var user model.User
	user.ID = userId
	err := r.db.QueryRow("SELECT first_name,last_name,email,birth_day,created_at FROM users WHERE id=$1", userId).Scan(&user.FirstName, &user.LastName, &user.Email, &user.BirthDay, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
