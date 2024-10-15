package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

func (br *BrandRepository) GetAllBrands(limit *int) ([]*models.Brand, error) {
	var brands []*models.Brand

	query := br.db
	if limit != nil {
		query = query.Limit(*limit)
	}

	if result := query.Find(&brands); result.Error != nil {
		return nil, result.Error
	}
	return brands, nil
}

func (br *BrandRepository) GetBrandByName(name string) (*models.Brand, error) {
	var brand *models.Brand
	if result := br.db.First(&brand, "name = ?", name); result.Error != nil {
		return nil, result.Error
	}
	return brand, nil
}

func (br *BrandRepository) GetBrandById(id uint) (*models.Brand, error) {
	var brand *models.Brand
	if result := br.db.
		Preload("Banners").
		First(&brand, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return brand, nil
}

func (br *BrandRepository) CreateBrand(brand *models.Brand) error {
	if result := br.db.Create(brand); result.Error != nil {
		return result.Error
	}
	return nil
}

func (br *BrandRepository) UpdateBrand(brand *models.Brand) error {
	if result := br.db.Save(brand); result.Error != nil {
		return result.Error
	}
	return nil
}

func (br *BrandRepository) DeleteBrand(id uint) error {
	if result := br.db.Delete(&models.Brand{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (br *BrandRepository) GetProductCountByCategoryAndBrand(categoryID, brandID uint) (int64, error) {
	var count int64

	if err := br.db.Model(&models.Products{}).
		Where("category_id = ? AND brand_id = ?", categoryID, brandID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (br *BrandRepository) GetCategoriesByBrand(brandID uint) ([]models.Category, error) {
	var categories []models.Category

	if err := br.db.Table("products").
		Select("categories.*").
		Joins("JOIN categories ON products.category_id = categories.id").
		Where("products.brand_id = ?", brandID).
		Group("categories.id").
		Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (br *BrandRepository) GetMainImage(categoryID uint, brandID uint) (string, error) {
	var mainImage string

	if err := br.db.Table("product_images").
		Select("image_url").
		Where("product_id IN (SELECT id FROM products WHERE category_id = ? AND brand_id = ?) AND is_main = true", categoryID, brandID).
		Limit(1).
		Scan(&mainImage).Error; err != nil {
		return "", err
	}

	return mainImage, nil
}
