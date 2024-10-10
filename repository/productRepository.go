package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"next_device/backend/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) CreateProduct(products *models.Products) error {
	if result := pr.db.Create(products); result.Error != nil {
		return result.Error
	}
	return nil
}
func (pr *ProductRepository) GetAllProduct() ([]*models.Products, error) {
	var products []*models.Products
	if result := pr.db.Find(&products); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return products, nil
		}
		return nil, result.Error
	}
	return products, nil
}

func (pr *ProductRepository) GetProductById(id uint) (*models.Products, error) {
	var product *models.Products
	if result := pr.db.
		Preload("Images").
		Preload("Variants").
		Preload("Brand").
		Preload("Category").
		First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (pr *ProductRepository) UpdateProduct(products *models.Products) error {
	if result := pr.db.Save(products); result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) DeleteProduct(id uint) error {
	if result := pr.db.Delete(&models.Products{ID: id}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) GetProductsByCategoryPaged(category string, limit, offset int) ([]*models.Products, error) {
	var products []*models.Products
	fmt.Print(category)

	if result := pr.db.
		Table("products").
		Select("products.id", "products.name", "products.description", "products.price").
		Joins("JOIN categories ON products.category_id = categories.id").
		Where("categories.name = ?", category).
		Limit(limit).
		Offset(offset).
		Find(&products); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return products, nil
		}
		return nil, result.Error
	}

	return products, nil
}

func (pr *ProductRepository) GetDiscountedProductsPaged(limit, offset int) ([]*models.Products, error) {
	var products []*models.Products

	if result := pr.db.
		Select("id", "name", "description", "price").
		Where("discount_price > ?", 0).
		Limit(limit).
		Offset(offset).
		Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (pr *ProductRepository) GetNewProductsPaged(limit, offset int) ([]*models.Products, error) {
	var products []*models.Products

	if result := pr.db.
		Select("id", "name", "description", "price").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
