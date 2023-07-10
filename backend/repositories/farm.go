package repositories

import (
	"fmt"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"gorm.io/gorm"
)

type IFarmRepository interface {
	All() ([]*models.Farm, error)
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
