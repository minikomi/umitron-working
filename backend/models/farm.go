package models

import (
	"gorm.io/gorm"
)

type Farm struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}
