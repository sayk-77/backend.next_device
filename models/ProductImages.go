package models

type ProductImage struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint   `json:"productId"`
	ImageURL  string `gorm:"not null" json:"imageUrl"`
	IsMain    bool   `json:"isMain"`

	Product Products `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
