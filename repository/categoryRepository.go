package repository

import (
	"gorm.io/gorm"
	"next_device/backend/models"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) GetAllCategory() ([]*models.Category, error) {
	var categories []*models.Category
	if err := cr.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (cr *CategoryRepository) GetCategoryById(id uint) (*models.Category, error) {
	var category models.Category
	if err := cr.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr *CategoryRepository) GetCategoryByName(name string) (*models.Category, error) {
	var category models.Category
	if err := cr.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (cr *CategoryRepository) CreateCategory(category *models.Category) error {
	if err := cr.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepository) UpdateCategory(category *models.Category) error {
	if err := cr.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepository) DeleteCategory(id uint) error {
	if err := cr.db.Delete(&models.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CategoryRepository) GetProductCountByCategory() (map[uint]int64, error) {
	var counts []struct {
		CategoryID   uint
		ProductCount int64
	}
	err := cr.db.
		Table("products").
		Select("category_id, COUNT(*) as product_count").
		Group("category_id").
		Find(&counts).Error
	if err != nil {
		return nil, err
	}

	result := make(map[uint]int64)
	for _, count := range counts {
		result[count.CategoryID] = count.ProductCount
	}
	return result, nil
}

func (cr *CategoryRepository) SearchCategory(query string) ([]*models.Category, error) {
	var categories []*models.Category

	if err := cr.db.
		Where("(title ILIKE ? OR SIMILARITY(title, ?) > 0.5)", "%"+query+"%", query).
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
