package services

import (
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/repositories"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -package=mocks -destination=../mocks/fish_pen_service.go
type IFishPenService interface {
	Create(farmID uint, fishPen *models.FishPen) (*models.FishPen, error)
	GetAllForFarm(farmID uint) ([]*models.FishPen, error)
	GetSingle(farmID, fishPenID uint) (*models.FishPen, error)
	Update(farmID uint, fishPen *models.FishPen) (*models.FishPen, error)
	Delete(farmID, fishPenID uint) error
}

type FishPenService struct {
	farmRepo    repositories.IFarmRepository
	fishpenRepo repositories.IFishPenRepository
}

func NewFishPenService(fr repositories.IFarmRepository, fpr repositories.IFishPenRepository) IFishPenService {
	return &FishPenService{farmRepo: fr, fishpenRepo: fpr}
}

func (s *FishPenService) Create(farmID uint, fishPen *models.FishPen) (*models.FishPen, error) {
	// Prevent duplicate naming within farm
	if exists, err := s.fishpenRepo.ExistsByNameForFarm(farmID, fishPen.Name); err != nil {
		return nil, err
	} else if exists {
		return nil, types.ErrDuplicate
	}

	fp, err := s.fishpenRepo.Create(fishPen)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

func (s *FishPenService) GetAllForFarm(farmID uint) ([]*models.FishPen, error) {
	if exists, err := s.farmRepo.ExistsByID(farmID); err != nil {
		return nil, err
	} else if !exists {
		return nil, types.ErrNotFound
	}

	fps, err := s.fishpenRepo.GetAllForFarm(farmID)
	if err != nil {
		return nil, err
	}
	return fps, nil
}

func (s *FishPenService) GetSingle(farmID, fishPenID uint) (*models.FishPen, error) {
	if exists, err := s.farmRepo.ExistsByID(farmID); err != nil {
		return nil, err
	} else if !exists {
		return nil, types.ErrNotFound
	}

	fp, err := s.fishpenRepo.GetByID(fishPenID)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

func (s *FishPenService) Update(farmID uint, fishPen *models.FishPen) (*models.FishPen, error) {
	if exists, err := s.farmRepo.ExistsByID(farmID); err != nil {
		return nil, err
	} else if !exists {
		return nil, types.ErrNotFound
	}

	// Prevent duplicate naming within farm
	if fp, err := s.fishpenRepo.GetByNameForFarm(farmID, fishPen.Name); err != nil {
		if err != types.ErrNotFound {
			return nil, err
		}
	} else if fp.ID != fishPen.ID {
		return nil, types.ErrDuplicate
	}

	fp, err := s.fishpenRepo.Update(fishPen)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

func (s *FishPenService) Delete(farmID, fishPenID uint) error {
	if exists, err := s.farmRepo.ExistsByID(farmID); err != nil {
		return err
	} else if !exists {
		return types.ErrNotFound
	}

	if err := s.fishpenRepo.DeleteByID(fishPenID); err != nil {
		return err
	}
	return nil
}
