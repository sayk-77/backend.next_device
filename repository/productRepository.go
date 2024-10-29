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

func (pr *ProductRepository) GetFilteredProducts(
	category string,
	priceFrom, priceTo *int,
	brands []string,
	screenFrom, screenTo *float64,
	memories, ram, ratings, cameraQualities, os []string,
	limit, offset int,
) ([]*models.Products, error) {

	// Вывод всех полученных фильтров
	if priceFrom != nil {
		fmt.Printf("Цена от: %d\n", *priceFrom)
	} else {
		fmt.Printf("Цена от: не указана\n")
	}

	if priceTo != nil {
		fmt.Printf("Цена до: %d\n", *priceTo)
	} else {
		fmt.Printf("Цена до: не указана\n")
	}

	fmt.Printf("Бренды: %v\n", brands)

	if screenFrom != nil {
		fmt.Printf("Экран от: %.2f\n", *screenFrom)
	} else {
		fmt.Printf("Экран от: не указан\n")
	}

	if screenTo != nil {
		fmt.Printf("Экран до: %.2f\n", *screenTo)
	} else {
		fmt.Printf("Экран до: не указан\n")
	}

	fmt.Printf("Память: %v\n", memories)
	fmt.Printf("Оперативная память: %v\n", ram)
	fmt.Printf("Рейтинг: %v\n", ratings)
	fmt.Printf("Качество камеры: %v\n", cameraQualities)
	fmt.Printf("ОС: %v\n", os)
	fmt.Printf("Лимит: %d, Смещение: %d\n", limit, offset)

	query := pr.db.Model(&models.Products{}).
		Joins("JOIN categories ON products.category_id = categories.id").
		Joins("JOIN brands ON products.brand_id = brands.id").
		Joins("JOIN product_filters ON products.id = product_filters.product_id").
		Where("categories.name = ?", category)

	if priceFrom != nil && priceTo != nil {
		query = query.Where("products.price BETWEEN ? AND ?", *priceFrom, *priceTo)
	} else if priceFrom != nil {
		query = query.Where("products.price >= ?", *priceFrom)
	} else if priceTo != nil {
		query = query.Where("products.price <= ?", *priceTo)
	}

	if len(brands) > 0 {
		query = query.Where("brands.name IN (?)", brands)
	}
	if screenFrom != nil {
		query = query.Where("product_filters.display_size >= ?", *screenFrom)
	}
	if screenTo != nil {
		query = query.Where("product_filters.display_size <= ?", *screenTo)
	}
	if len(memories) > 0 {
		query = query.Where("product_filters.storage IN (?)", memories)
	}
	if len(ram) > 0 {
		query = query.Where("product_filters.ram IN (?)", ram)
	}
	if len(ratings) > 0 {
		query = query.Where("products.rating IN (?)", ratings)
	}
	if len(cameraQualities) > 0 {
		query = query.Where("product_filters.camera_quality IN (?)", cameraQualities)
	}
	if len(os) > 0 {
		query = query.Where("product_filters.os IN (?)", os)
	}

	query = query.Limit(limit).Offset(offset)

	var products []*models.Products
	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	// Проверка наличия продуктов
	if len(products) == 0 {
		fmt.Println("Не найдено продуктов с указанными фильтрами.")
	} else {
		fmt.Printf("Найденные продукты: %+v\n", products)
	}

	return products, nil
}
