package database

import (
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"gorm.io/gorm"
)

func DropDB(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&models.Farm{},
		&models.FishMaker{},
		&models.FishBreed{},
		&models.Juvenile{},
		&models.FishPen{},
		&models.FishLot{},
	)
	if err != nil {
		return err
	}

	return nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Farm{},
		&models.FishMaker{},
		&models.FishBreed{},
		&models.Juvenile{},
		&models.FishPen{},
		&models.FishLot{},
	)
	if err != nil {
		return err
	}

	return nil
}
