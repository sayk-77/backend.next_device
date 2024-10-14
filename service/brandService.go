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

func (bs *BrandService) GetAllBrand(limit *int) ([]*models.Brand, error) {
	return bs.brandRepo.GetAllBrands(limit)
}

func (bs *BrandService) GetBrandByName(name string) (*models.Brand, error) {
	return bs.brandRepo.GetBrandByName(name)
}

func (bs *BrandService) GetBrandByID(id uint) (*models.Brand, error) {
	return bs.brandRepo.GetBrandById(id)
}

func (bs *BrandService) CreateBrand(brand *models.Brand) error {
	return bs.brandRepo.CreateBrand(brand)
}

func (bs *BrandService) UpdateBrand(brand *models.Brand) error {
	return bs.brandRepo.UpdateBrand(brand)
}

func (bs *BrandService) DeleteBrand(id uint) error {
	return bs.brandRepo.DeleteBrand(id)
}

func (bs *BrandService) GetCategoriesByBrand(brandID uint) ([]models.CategoryWithCountAndImage, error) {
	categories, err := bs.brandRepo.GetCategoriesByBrand(brandID)
	if err != nil {
		return nil, err
	}

	var categoriesWithCount []models.CategoryWithCountAndImage

	for _, category := range categories {
		count, err := bs.brandRepo.GetProductCountByCategory(category.ID)
		if err != nil {
			return nil, err
		}

		mainImage, err := bs.brandRepo.GetMainImage(category.ID, brandID)
		if err != nil {
			return nil, err
		}

		categoriesWithCount = append(categoriesWithCount, models.CategoryWithCountAndImage{
			Category:      category,
			Count:         count,
			ImageCategory: mainImage,
		})
	}

	return categoriesWithCount, nil
}
