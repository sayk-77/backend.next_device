package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (cs *CategoryService) GetAllCategory() ([]*models.Category, error) {
	return cs.categoryRepo.GetAllCategory()
}

func (cs *CategoryService) GetCategoryById(id uint) (*models.Category, error) {
	return cs.categoryRepo.GetCategoryById(id)
}

func (cs *CategoryService) GetCategoryByName(name string) (*models.Category, error) {
	return cs.categoryRepo.GetCategoryByName(name)
}

func (cs *CategoryService) CreateCategory(category *models.Category) error {
	return cs.categoryRepo.CreateCategory(category)
}

func (cs *CategoryService) UpdateCategory(category *models.Category) error {
	return cs.categoryRepo.UpdateCategory(category)
}

func (cs *CategoryService) DeleteCategoryById(id uint) error {
	return cs.categoryRepo.DeleteCategory(id)
}

func (cs *CategoryService) GetAllCategoryAndCount() ([]*models.CategoryCount, error) {

	categories, err := cs.categoryRepo.GetAllCategory()
	if err != nil {
		return nil, err
	}

	productCounts, err := cs.categoryRepo.GetProductCountByCategory()
	if err != nil {
		return nil, err
	}

	var result []*models.CategoryCount
	for _, category := range categories {
		count := productCounts[category.ID]
		result = append(result, &models.CategoryCount{
			Category: *category,
			Count:    int(count),
		})
	}

	return result, nil
}
