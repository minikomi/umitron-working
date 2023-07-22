package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -package=mocks -destination=../mocks/farm_repository.go
type IFarmRepository interface {
	All() ([]*models.Farm, error)
	GetByID(id uint) (*models.Farm, error)
	ExistsByID(id uint) (bool, error)
}

type FarmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) IFarmRepository {
	return &FarmRepository{
		db: db,
	}
}

func (r *FarmRepository) All() ([]*models.Farm, error) {
	var fs []*models.Farm

	if err := r.db.Find(&fs).Error; err != nil {
		return nil, fmt.Errorf("failed to find farms: %w", err)
	}

	return fs, nil
}

func (r *FarmRepository) GetByID(id uint) (*models.Farm, error) {
	var f models.Farm

	if err := r.db.First(&f, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, types.ErrNotFound
		}
		return nil, fmt.Errorf("failed to find farm by id %d: %w", id, err)
	}

	return &f, nil
}

func (r *FarmRepository) ExistsByID(farmID uint) (bool, error) {
	var farm models.Farm
	if err := r.db.First(&farm, farmID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("failed to find farm by id %d: %w", farmID, err)
	}
	return true, nil
}
