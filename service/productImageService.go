package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type ProductImageService struct {
	imageRepo *repository.ProductImageRepository
}

func NewProductImageService(imageRepo *repository.ProductImageRepository) *ProductImageService {
	return &ProductImageService{imageRepo: imageRepo}
}

func (pis *ProductImageService) GetMainImage(productId uint) (*models.ProductImage, error) {
	return pis.imageRepo.GetMainImage(productId)
}
