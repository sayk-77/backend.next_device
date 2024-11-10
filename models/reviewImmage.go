package models

type ReviewImage struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ReviewId uint   `gorm:"foreignKey;autoIncrement;index" json:"reviewId"`
	ImageUrl string `gorm:"size:255;" json:"imageUrl"`
}
