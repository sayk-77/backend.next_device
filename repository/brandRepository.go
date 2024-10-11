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

func (br *BrandRepository) GetAllBrands() ([]*models.Brand, error) {
	var brands []*models.Brand
	if result := br.db.Find(&brands); result.Error != nil {
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
	if result := br.db.First(&brand, "id = ?", id); result.Error != nil {
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
