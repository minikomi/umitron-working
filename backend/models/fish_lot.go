package models

import (
	"time"

	"gorm.io/gorm"
)

type FishLot struct {
	gorm.Model
	FarmId       uint `gorm:"not null"`
	Farm         Farm
	JuvenileId   uint `gorm:"not null"`
	Juvenile     Juvenile
	FishPenId    uint `gorm:"not null"`
	FishPen      FishPen
	Name         string    `gorm:"type:varchar(255);not null"`
	InitialCount uint      `gorm:"not null"`
	StartAt      time.Time `gorm:"not null"`
	EndAt        time.Time `gorm:"not null"`
}
