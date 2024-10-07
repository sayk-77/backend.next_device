package models

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `json:"userId"`
	ProductID uint      `json:"productId"`
	Rating    int       `gorm:"not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Review    string    `gorm:"type:text" json:"review"`
	CreatedAt time.Time `json:"createdAt"`

	User    User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Product Products `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
