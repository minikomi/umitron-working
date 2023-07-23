package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -package=mocks -destination=../mocks/fish_pen_repository.go
type IFishPenRepository interface {
	Create(fishPen *models.FishPen) (*models.FishPen, error)
	GetAllForFarm(farmID uint) ([]*models.FishPen, error)
	GetByID(fishPenID uint) (*models.FishPen, error)
	GetByNameForFarm(farmID uint, name string) (*models.FishPen, error)
	ExistsByID(fishPenID uint) (bool, error)
	ExistsByNameForFarm(farmID uint, name string) (bool, error)
	Update(fishPen *models.FishPen) (*models.FishPen, error)
	DeleteByID(fishPenID uint) error
}

type FishPenRepository struct {
	db *gorm.DB
}

func NewFishPenRepository(db *gorm.DB) IFishPenRepository {
	return &FishPenRepository{db: db}
}

func (r *FishPenRepository) Create(fishPen *models.FishPen) (*models.FishPen, error) {
	if err := r.db.Create(fishPen).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, types.ErrInvalid
		}

		return nil, fmt.Errorf("failed to create fish pen: %w", err)
	}

	return fishPen, nil
}

func (r *FishPenRepository) GetAllForFarm(farmID uint) ([]*models.FishPen, error) {
	var fps []*models.FishPen

	if err := r.db.Where("farm_id = ?", farmID).Find(&fps).Error; err != nil {
		return nil, fmt.Errorf("failed to find fish pens for farm %d: %w", farmID, err)
	}

	return fps, nil
}

func (r *FishPenRepository) GetByID(fishPenID uint) (*models.FishPen, error) {
	var fp models.FishPen

	if err := r.db.Where("id = ?", fishPenID).First(&fp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, types.ErrNotFound
		} else if errors.Is(err, gorm.ErrInvalidData) {
			return nil, types.ErrInvalid
		}
		return nil, fmt.Errorf("failed to find fish pen by id %d: %w", fishPenID, err)
	}

	return &fp, nil
}

func (r *FishPenRepository) GetByNameForFarm(farmID uint, name string) (*models.FishPen, error) {
	var fp models.FishPen

	if err := r.db.Where("farm_id = ? AND name = ?", farmID, name).First(&fp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, types.ErrNotFound
		} else if errors.Is(err, gorm.ErrInvalidData) {
			return nil, types.ErrInvalid
		}
		return nil, fmt.Errorf("failed to find fish pen for farm %d by name %s: %w", farmID, name, err)
	}

	return &fp, nil
}

func (r *FishPenRepository) ExistsByID(fishPenID uint) (bool, error) {
	err := r.db.Where("id = ?", fishPenID).First(&models.FishPen{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *FishPenRepository) ExistsByNameForFarm(farmID uint, name string) (bool, error) {
	err := r.db.Where("farm_id = ? AND name = ?", farmID, name).First(&models.FishPen{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *FishPenRepository) Update(fishPen *models.FishPen) (*models.FishPen, error) {
	updatableFields := map[string]interface{}{
		"Name":           fishPen.Name,
		"MakerModelName": fishPen.MakerModelName,
		"Description":    fishPen.Description,
		"Material":       fishPen.Material,
		"NetMaterial":    fishPen.NetMaterial,
		"Category":       fishPen.Category,
		"WidthCM":        fishPen.WidthCM,
		"LengthCM":       fishPen.LengthCM,
		"HeightCM":       fishPen.HeightCM,
	}

	if err := r.db.Model(fishPen).Updates(updatableFields).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, types.ErrInvalid
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, types.ErrNotFound
		}
		return nil, err
	}
	return fishPen, nil
}

func (r *FishPenRepository) DeleteByID(fishPenID uint) error {
	if err := r.db.Delete(&models.FishPen{}, fishPenID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return types.ErrNotFound
		}
		return err
	}
	return nil
}
