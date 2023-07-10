package database

import (
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) error {
	farm1 := &models.Farm{Name: "Farm 1"}
	if err := db.Where(farm1).FirstOrCreate(farm1).Error; err != nil {
		return err
	}

	farm2 := &models.Farm{Name: "Farm 2"}
	if err := db.Where(farm2).FirstOrCreate(farm2).Error; err != nil {
		return err
	}

	return nil
}
