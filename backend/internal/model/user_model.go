package model

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	BirthDay  time.Time `json:"birth_day" binding:"required" time_format:"2006-01-02"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserInput struct {
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	BirthDay  time.Time `json:"birth_day" binding:"required"`
}
type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
