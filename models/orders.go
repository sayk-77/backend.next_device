package models

import (
	"gorm.io/gorm"
	"next_device/backend/types"
)

type Orders struct {
	gorm.Model
	UserId     int              `json:"userId"`
	TotalPrice float64          `json:"totalPrice"`
	Status     types.StatusType `json:"status"`
}
