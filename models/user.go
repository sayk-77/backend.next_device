package models

import (
	"gorm.io/gorm"
	"next_device/backend/types"
)

type User struct {
	gorm.Model
	Email        string               `json:"email"`
	PasswordHash string               `json:"passwordHash"`
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastName"`
	Role         types.RoleCustomType `json:"role"`
}
