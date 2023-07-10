package database

import (
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Farm{},
	)
	if err != nil {
		return err
	}

	return nil
}
