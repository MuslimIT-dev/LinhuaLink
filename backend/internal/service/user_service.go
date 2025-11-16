package service

import (
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/repository"
	"LinhuaLink/backend/pkg/utils"
)

type UserService interface {
	SignUp(input model.RegisterUserInput) (model.User, error)
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) SignUp(json model.RegisterUserInput) (model.User, error) {
	json.Password = utils.GenerateHashFromPassword(json.Password) // hash password
	user, err := s.repo.Signup(json.FirstName, json.LastName, json.Email, json.Password, json.BirthDay)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
