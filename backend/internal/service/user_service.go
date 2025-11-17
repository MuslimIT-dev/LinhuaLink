package service

import (
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/repository"
	"LinhuaLink/backend/pkg/utils"
	"errors"
)

type UserService interface {
	SignUp(input model.RegisterUserInput) (model.User, error)
	Login(json model.LoginUserInput) (model.User, error)
	Me(userId int) (model.User, error)
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

func (s *userService) Login(json model.LoginUserInput) (model.User, error) {
	user, password, err := s.repo.Login(json.Email)
	if err != nil {
		return model.User{}, err
	}
	if utils.ComparePasswords(password, json.Password) {
		return user, nil
	}
	return model.User{}, errors.New("password is incorrect")
}

func (s *userService) Me(userId int) (model.User, error) {
	user, err := s.repo.Me(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
