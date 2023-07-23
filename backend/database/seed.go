package database

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/models"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/types"
)

func SeedDB(db *gorm.DB) error {
	// Seed Farms
	fmt.Println("Seeding Farms...")

	farm1 := &models.Farm{Name: "Farm 1"}
	if err := db.Where(farm1).FirstOrCreate(farm1).Error; err != nil {
		return err
	}

	farm2 := &models.Farm{Name: "Farm 2"}
	if err := db.Where(farm2).FirstOrCreate(farm2).Error; err != nil {
		return err
	}

	fmt.Println("Seeding Fish Pens farm 1")

	fishPenNames := []string{"A1", "A2", "A3", "B3", "C2", "D4", "E5", "F6", "XXXX", "Other"}

	for idx, name := range fishPenNames {
		description := fmt.Sprintf("Fish Pen %s description", name)
		material := fmt.Sprintf("Fish Pen %s Material", name)
		netMaterial := fmt.Sprintf("Fish Pen %s Net Material", name)
		makerModelName := fmt.Sprintf("FISH PEN %s", name)
		fishPen := &models.FishPen{
			FarmID:         farm1.ID,
			Name:           name,
			MakerModelName: &makerModelName,
			Description:    &description,
			Material:       &material,
			NetMaterial:    &netMaterial,
			Category: []types.FishPenCategory{
				types.FishPenCategoryFloating,
				types.FishPenCategoryFixed,
				types.FishPenCategorySubmersible,
				types.FishPenCategorySubmersed,
				types.FishPenCategoryOther,
			}[idx%5],
			WidthCM:  uint(100 * (idx%3 + 1)),
			LengthCM: uint(100 * (idx%5 + 1)),
			HeightCM: uint(100 * (idx%7 + 1)),
		}
		if err := db.Where(fishPen).FirstOrCreate(fishPen).Error; err != nil {
			return err
		}
	}

	return nil
}
