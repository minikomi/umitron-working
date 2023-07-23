package models

import (
	"time"

	"gorm.io/gorm"
)

type Juvenile struct {
	gorm.Model
	FarmId            uint `gorm:"not null"`
	Farm              Farm
	FishMakerId       uint `gorm:"not null"`
	FishMaker         FishMaker
	FishBreedId       uint `gorm:"not null"`
	FishBreed         FishBreed
	Name              string `gorm:"type:varchar(255);not null"`
	Description       *string
	InitialFishSizeMM uint      `gorm:"type:int;not null"`
	PriceJPY          uint      `gorm:"type:int;not null"`
	Count             uint      `gorm:"type:int;not null"`
	PurchasedAt       time.Time `gorm:"type:date;not null"`
	FishLots          []FishLot
}
