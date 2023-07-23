package services

import (
	"fmt"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/repositories"
	"gorm.io/gorm"
)

var (
	FarmNotFoundErr = fmt.Errorf("farm not found")
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -package=mocks -destination=../mocks/farm_service.go
type IFarmService interface {
	GetAll() ([]*models.Farm, error)
	GetByID(id uint) (*models.Farm, error)
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

func (s *FarmService) GetByID(id uint) (*models.Farm, error) {
	f, err := s.farmRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("%w: %d", FarmNotFoundErr, id)
		}

		return nil, err
	}
	return f, nil
}
