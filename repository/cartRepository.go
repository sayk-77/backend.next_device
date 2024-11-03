package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddItem(cartItem *models.CartItem) error {
	return r.db.Create(cartItem).Error
}

func (r *CartRepository) RemoveItem(cartID, productID uint) error {
	return r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&models.CartItem{}).Error
}

func (r *CartRepository) GetCartItems(userID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	var cart models.Cart

	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return items, err
	}

	err = r.db.Where("cart_id = ?", cart.ID).Find(&items).Error
	if err != nil {
		return items, err
	}

	return items, err
}

func (r *CartRepository) GetCartByUserID(userID uint) (uint, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	return cart.ID, err
}

func (r *CartRepository) CreateCart(cart *models.Cart) (uint, error) {
	err := r.db.Create(cart).Error
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}
