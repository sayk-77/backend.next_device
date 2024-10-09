package repository

import (
	"errors"
	"gorm.io/gorm"
	"next_device/backend/models"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (cr *CartRepository) CreateCart(cart *models.Cart) error {
	if result := cr.db.Create(&cart); result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CartRepository) GetCartById(id uint) (*models.Cart, error) {
	var cart *models.Cart
	if result := cr.db.First(&cart, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return cart, result.Error
		}
		return nil, result.Error
	}
	return cart, nil
}

func (cr *CartRepository) GetCartByUserId(id uint) (*models.Cart, error) {
	var cart *models.Cart
	if result := cr.db.First(&cart, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return cart, result.Error
		}
		return nil, result.Error
	}
	return cart, nil
}
