package models

type ProductVariant struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID    uint   `json:"productId"`
	VariantName  string `gorm:"not null" json:"variantName"`
	VariantValue string `gorm:"not null" json:"variantValue"`

	Product Products `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
