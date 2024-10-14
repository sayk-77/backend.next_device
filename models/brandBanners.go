package models

type BrandBanners struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	BrandId  uint   `json:"brandId"`
	Title    string `json:"title"`
	ImageUrl string `json:"imageUrl"`
}
