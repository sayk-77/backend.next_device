package repository

import (
	"errors"
	"gorm.io/gorm"
	"next_device/backend/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (or *OrderRepository) CreateOrder(order *models.Order) error {
	if err := or.db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) GetOrderById(id uint) (*models.Order, error) {
	var order *models.Order
	if result := or.db.First(&order, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return order, nil
}

func (or *OrderRepository) GetAllOrders() ([]*models.Order, error) {
	var orders []*models.Order
	if result := or.db.Find(&orders); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return orders, nil
		}
		return nil, result.Error
	}

	return orders, nil
}

func (or *OrderRepository) UpdateOrder(order *models.Order) error {
	if result := or.db.Save(*order); result.Error != nil {
		return result.Error
	}
	return nil
}

func (or *OrderRepository) DeleteOrder(id uint) error {
	if result := or.db.Delete(&models.Order{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}
