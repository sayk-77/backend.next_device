package repository

import (
	"errors"
	"gorm.io/gorm"
	"next_device/backend/models"
)

type ProductRepositoru struct {
	db *gorm.DB
}

func NewProductRepositoru(db *gorm.DB) *ProductRepositoru {
	return &ProductRepositoru{db: db}
}

func (pr *ProductRepositoru) CreateProduct(products *models.Products) error {
	if result := pr.db.Create(products); result.Error != nil {
		return result.Error
	}
	return nil
}
func (pr *ProductRepositoru) GetAllProduct() ([]*models.Products, error) {
	var products []*models.Products
	if result := pr.db.Find(&products); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return products, nil
		}
		return nil, result.Error
	}
	return products, nil
}

func (pr *ProductRepositoru) GetProductById(id uint) (*models.Products, error) {
	var product *models.Products
	if result := pr.db.First(&product, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return product, nil
}

func (pr *ProductRepositoru) UpdateProduct(products *models.Products) error {
	if result := pr.db.Save(products); result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepositoru) DeleteProduct(id uint) error {
	if result := pr.db.Delete(&models.Products{ID: id}); result.Error != nil {
		return result.Error
	}
	return nil
}
