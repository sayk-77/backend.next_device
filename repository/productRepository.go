package repository

import (
	"errors"
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
		Preload("Details").
		First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (pr *ProductRepository) GetProductByName(name string) (*models.Products, error) {
	var product *models.Products
	if result := pr.db.
		Preload("Images").
		Preload("Variants").
		Preload("Brand").
		Preload("Category").
		Preload("Details").
		Where("search_name = ?", name).
		First(&product); result.Error != nil {
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

func (pr *ProductRepository) GetProductsByCategoryPaged(category string, limit, offset int) ([]models.Products, string, error) {
	var products []models.Products
	var categoryTitle string

	if result := pr.db.
		Table("products").
		Select("products.*, categories.title AS category_title").
		Joins("JOIN categories ON products.category_id = categories.id").
		Where("categories.name = ?", category).
		Limit(limit).
		Offset(offset).
		Scan(&products); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return products, "", nil
		}
		return nil, "", result.Error
	}

	if err := pr.db.
		Table("categories").
		Select("title").
		Where("name = ?", category).
		Scan(&categoryTitle).Error; err != nil {
		return nil, "", err
	}

	return products, categoryTitle, nil
}

func (pr *ProductRepository) GetDiscountedProductsPaged(limit, offset int, brand string) ([]*models.Products, error) {
	var products []*models.Products

	query := pr.db.Select("products.id, products.name, products.description, products.price, products.discount_price, products.search_name").
		Joins("JOIN brands ON brands.id = products.brand_id").
		Where("products.discount_price > ?", 0).
		Order("products.discount_price DESC").
		Limit(limit).
		Offset(offset)

	if brand != "" {
		query = query.Where("brands.name ILIKE ?", "%"+brand+"%")
	}

	if result := query.Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (pr *ProductRepository) GetNewProductsPaged(limit, offset int) ([]*models.Products, error) {
	var products []*models.Products

	if result := pr.db.
		Select("id", "name", "description", "price", "search_name").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (pr *ProductRepository) GetProductsByBrandAndCategory(brandId uint, categoryName string, limit, offset int) ([]*models.Products, error) {
	var products []*models.Products

	if result := pr.db.
		Joins("JOIN categories ON products.category_id = categories.id").
		Where("products.brand_id = ? AND categories.title = ?", brandId, categoryName).
		Limit(limit).Offset(offset).
		Find(&products); result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (pr *ProductRepository) SearchProduct(query string, limit, offset int) ([]*models.Products, error) {
	var products []*models.Products

	if result := pr.db.
		Where("(name ILIKE ? OR SIMILARITY(name, ?) > 0.5)", "%"+query+"%", query).
		Limit(limit).Offset(offset).
		Find(&products); result.Error != nil {
		return nil, result.Error
	}
	if len(products) == 0 {
		return nil, nil
	}
	return products, nil
}
