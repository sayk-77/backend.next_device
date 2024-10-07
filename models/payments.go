package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID       uint      `json:"orderId"`
	PaymentMethod string    `gorm:"type:enum('credit_card', 'paypal', 'bank_transfer');not null" json:"paymentMethod"`
	PaymentStatus string    `gorm:"type:enum('pending', 'completed', 'failed');default:'pending';not null" json:"paymentStatus"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentDate   time.Time `json:"paymentDate"`

	Order Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
