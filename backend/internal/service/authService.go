package service

import (
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/repository"
	"LinhuaLink/backend/pkg/utils"
	"errors"
)

type AuthService interface {
	SignUp(input model.RegisterUserInput) (model.User, error)
	Login(json model.LoginUserInput) (model.User, error)
}
type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SignUp(json model.RegisterUserInput) (model.User, error) {
	json.Password = utils.GenerateHashFromPassword(json.Password) // hash password
	user, err := s.repo.Signup(json.FirstName, json.LastName, json.Email, json.Password, json.BirthDay)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *authService) Login(json model.LoginUserInput) (model.User, error) {
	user, password, err := s.repo.Login(json.Email)
	if err != nil {
		return model.User{}, err
	}
	if utils.ComparePasswords(password, json.Password) {
		return user, nil
	}
	return model.User{}, errors.New("password is incorrect")
}
