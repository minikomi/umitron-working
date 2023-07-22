package models

import (
	"gorm.io/gorm"
)

type FishBreed struct {
	gorm.Model
	Name              string `gorm:"type:varchar(255);not null"`
	ScientificName    string `gorm:"type:varchar(255);not null"`
	FeedType          *string
	Description       *string
	TimeToRaiseMonths *uint
}
