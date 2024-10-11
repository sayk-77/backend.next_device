package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type BrandService struct {
	brandRepo *repository.BrandRepository
}

func NewBrandBrand(brandRepo *repository.BrandRepository) *BrandService {
	return &BrandService{brandRepo: brandRepo}
}

func (br *BrandService) GetAllBrand() ([]*models.Brand, error) {
	return br.brandRepo.GetAllBrands()
}

func (br *BrandService) GetBrandByName(name string) (*models.Brand, error) {
	return br.brandRepo.GetBrandByName(name)
}

func (br *BrandService) GetBrandByID(id uint) (*models.Brand, error) {
	return br.brandRepo.GetBrandById(id)
}

func (br *BrandService) CreateBrand(brand *models.Brand) error {
	return br.brandRepo.CreateBrand(brand)
}

func (br *BrandService) UpdateBrand(brand *models.Brand) error {
	return br.brandRepo.UpdateBrand(brand)
}

func (br *BrandService) DeleteBrand(id uint) error {
	return br.brandRepo.DeleteBrand(id)
}
