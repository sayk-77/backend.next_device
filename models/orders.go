package models

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint      `json:"userId"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null" json:"totalPrice"`
	Status     string    `gorm:"type:order_status;default:'pending';not null" json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`

	User       User        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orderItems"`
	Payment    Payment     `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"payment"`
}
