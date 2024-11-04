package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (repo *PaymentRepository) SavePayment(payment *models.Payment) error {
	return repo.db.Create(payment).Error
}
