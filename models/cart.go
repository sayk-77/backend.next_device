package models

type Cart struct {
	ID        uint  `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint  `json:"userId"`
	ProductID uint  `json:"productId"`
	VariantID *uint `json:"variantId"`
	Quantity  int   `gorm:"not null" json:"quantity"`

	User    User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Product Products       `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Variant ProductVariant `gorm:"foreignKey:VariantID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"variant"`
}
