package models

type Brand struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ImageUrl string `json:"imageUrl"`
	Name     string `gorm:"not null" json:"name"`

	Banners []BrandBanners `gorm:"foreignKey:BrandId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"banners"`
}
