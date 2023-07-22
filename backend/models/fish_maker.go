package models

import "gorm.io/gorm"

type FishMaker struct {
	gorm.Model
	FarmId       uint `gorm:"not null"`
	Farm         Farm
	Name         string `gorm:"type:varchar(255);not null"`
	Abbreviation string `gorm:"type:varchar(255);not null"`
	Address      *string
	PhoneNumber  *string
	EmailAddress *string
	FishBreeds   []FishBreed `gorm:"many2many:fishmaker_fishbreeds;"`
}
