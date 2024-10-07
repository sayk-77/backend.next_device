package models

type Category struct {
	ID               uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string `gorm:"not null" json:"name"`
	ParentCategoryID *uint  `json:"parentCategoryId"`

	SubCategories []Category `gorm:"foreignKey:ParentCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subCategories"`
}
