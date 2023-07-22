package models

import (
	"gorm.io/gorm"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

type FishPen struct {
	gorm.Model
	FarmID         uint `gorm:"not null;index:idx_farm_id_id,idx_farm_id_name"`
	Farm           Farm
	Name           string  `gorm:"type:varchar(255);not null;index:idx_farm_id_name"`
	MakerModelName *string `gorm:"type:varchar(255)"`
	Description    *string
	Material       *string
	NetMaterial    *string
	Category       types.FishPenCategory `gorm:"type:enum('fixed','floating','submersible','submersed','other');not null"`
	WidthCM        uint                  `gorm:"type:int"`
	LengthCM       uint                  `gorm:"type:int"`
	HeightCM       uint                  `gorm:"type:int"`
}
