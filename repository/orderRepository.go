package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindOrderById(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Where("id = ?", id).
		Preload("OrderItems.Products.Images").
		Preload("Payment").
		Preload("Address").
		First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) CreateOrderItems(items []models.OrderItem) error {
	return r.db.Create(&items).Error
}

func (r *OrderRepository) DeleteOrder(orderId uint) error {
	if err := r.db.Where("order_id = ?", orderId).Delete(&models.OrderItem{}).Error; err != nil {
		return err
	}
	return r.db.Delete(&models.Order{}, "id = ?", orderId).Error
}
