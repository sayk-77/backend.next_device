package models

type Category struct {
	ID               uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string     `gorm:"not null" json:"name"`
	CategoryImage    string     `json:"categoryImage"`
	ParentCategoryID *uint      `json:"parentCategoryId"`
	Title            string     `json:"title"`
	SubCategories    []Category `gorm:"foreignKey:ParentCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subCategories"`
}
