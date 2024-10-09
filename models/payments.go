package models

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID       uint      `json:"orderId"`
	PaymentMethod string    `gorm:"type:payment_method;not null" json:"paymentMethod"`
	PaymentStatus string    `gorm:"type:payment_status;default:'pending';not null" json:"paymentStatus"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentDate   time.Time `json:"paymentDate"`
}
