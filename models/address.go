package models

type Address struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint   `json:"userId"`
	Country    string `gorm:"not null" json:"country"`
	City       string `gorm:"not null" json:"city"`
	Street     string `gorm:"not null" json:"street"`
	PostalCode string `gorm:"not null" json:"postalCode"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
