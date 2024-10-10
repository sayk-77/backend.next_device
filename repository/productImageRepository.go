package repository

import (
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
		return nil, result.Error
	}

	return mainImage, nil
}
