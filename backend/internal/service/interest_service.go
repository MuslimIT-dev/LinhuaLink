package service

import (
	"LinhuaLink/backend/internal/model"
	"LinhuaLink/backend/internal/repository"
)

type InterestService interface {
	GetInterest() (map[string][]model.Interest, error)
}
type interestService struct {
	interestRepo repository.InterestRepository
}

func NewInterestService(interestRepo repository.InterestRepository) InterestService {
	return &interestService{interestRepo: interestRepo}
}

func (s *interestService) GetInterest() (map[string][]model.Interest, error) {
	return s.interestRepo.GetInterest()
}
