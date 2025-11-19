package model

import "time"

type User struct {
	ID        int       `json:"id" example:"1"`
	FirstName string    `json:"first_name" example:"Jon"`
	LastName  string    `json:"last_name" example:"Smith"`
	Email     string    `json:"email" example:"test@mail.com"`
	BirthDay  time.Time `json:"birth_day" example:"2006-01-02T15:04:05Z07:00"`
	CreatedAt time.Time `json:"created_at" example:"2025-11-18T21:00:00Z"`
}

type RegisterUserInput struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDay  time.Time `json:"birth_day"`
}
type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
