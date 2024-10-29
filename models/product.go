package models

import "time"

type Products struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SKU           string    `gorm:"uniqueIndex;not null" json:"sku"`
	Name          string    `gorm:"not null" json:"name"`
	SearchName    string    `json:"searchName"`
	Description   string    `gorm:"type:text" json:"description"`
	Price         float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	DiscountPrice float64   `gorm:"type:decimal(10,2);default=0'" json:"discountPrice"`
	Stock         int       `gorm:"not null" json:"stock"`
	CategoryID    uint      `json:"categoryId"`
	BrandID       uint      `json:"brandId"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`

	Category Category         `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	Brand    Brand            `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"brand"`
	Variants []ProductVariant `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"variants"`
	Images   []ProductImage   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images"`
	Reviews  []Review         `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"`
	Details  ProductDetails   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"details"`
	Filters  []ProductFilter  `gorm:"foreignKey:ProductID"`
}
