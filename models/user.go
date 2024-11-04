package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"password"`
	FirstName    string    `gorm:"not null" json:"firstName"`
	LastName     string    `gorm:"not null" json:"lastName"`
	Role         string    `gorm:"type:role;default:'customer';not null" json:"role"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`

	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"addresses"`
	Orders    []Order   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
	Reviews   []Review  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews"`
	CartItems []Cart    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cartItems"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}

type UserProfileResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Addresses []Address `json:"addresses"`
	Orders    []Order   `json:"orders"`
	Reviews   []Review  `json:"reviews"`
	CartItems []Cart    `json:"cartItems"`
}
