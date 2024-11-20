package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"next_device/backend/models"
)

type ProductImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) *ProductImageRepository {
	return &ProductImageRepository{db: db}
}

func (pir *ProductImageRepository) GetMainImage(productId uint) (*models.ProductImage, error) {
	var mainImage *models.ProductImage

	if result := pir.db.
		Where("product_id = ? AND is_main = ?", productId, true).
		First(&mainImage); result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("main image for product with ID %d not found", productId)
		}
		return nil, result.Error
	}

	return mainImage, nil
}
