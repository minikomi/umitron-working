package services

import (
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/repositories"
)

type IFarmService interface {
	GetAll() ([]*models.Farm, error)
}

type FarmService struct {
	farmRepo repositories.IFarmRepository
}

func NewFarmService(r repositories.IFarmRepository) IFarmService {
	return &FarmService{farmRepo: r}
}

func (s *FarmService) GetAll() ([]*models.Farm, error) {
	fs, err := s.farmRepo.All()
	if err != nil {
		return nil, err
	}
	return fs, nil
}
