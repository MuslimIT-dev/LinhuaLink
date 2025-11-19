package service

import (
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/repository"
)

type UserService interface {
	Me(userId int) (model.User, error)
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Me(userId int) (model.User, error) {
	user, err := s.repo.Me(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
