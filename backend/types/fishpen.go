package types

type FishPenCategory string

const (
	FishPenCategoryFixed       FishPenCategory = "fixed"
	FishPenCategoryFloating    FishPenCategory = "floating"
	FishPenCategorySubmersible FishPenCategory = "submersible"
	FishPenCategorySubmersed   FishPenCategory = "submersed"
	FishPenCategoryOther       FishPenCategory = "other"
)
