package models

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	ID        uint
	Name      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
