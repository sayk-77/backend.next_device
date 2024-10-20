package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type ProductDetailsRepository struct {
	db *gorm.DB
}

func NewProductDetailsRepository(db *gorm.DB) *ProductDetailsRepository {
	return &ProductDetailsRepository{db: db}
}

func (pdr *ProductDetailsRepository) GetProductDetailsByProductId(productId uint) (*models.ProductDetails, error) {
	var productDetails models.ProductDetails
	if err := pdr.db.Where("product_id = ?", productId).First(&productDetails).Error; err != nil {
		return nil, err
	}
	return &productDetails, nil
}

func (pdr *ProductDetailsRepository) CreateProductDetails(productDetails *models.ProductDetails) error {
	if err := pdr.db.Create(&productDetails).Error; err != nil {
		return err
	}
	return nil
}

func (pdr *ProductDetailsRepository) UpdateProductDetails(productDetails *models.ProductDetails) error {
	if err := pdr.db.Save(&productDetails).Error; err != nil {
		return err
	}
	return nil
}

func (pdr *ProductDetailsRepository) DeleteProductDetails(productId uint) error {
	if err := pdr.db.Delete(models.ProductDetails{ID: productId}).Error; err != nil {
		return err
	}
	return nil
}
