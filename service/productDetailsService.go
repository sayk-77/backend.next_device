package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type ProductDetailsService struct {
	productDetailsRep *repository.ProductDetailsRepository
}

func NewProductDetailsService(productDetailsRep *repository.ProductDetailsRepository) *ProductDetailsService {
	return &ProductDetailsService{productDetailsRep: productDetailsRep}
}

func (pds *ProductDetailsService) GetProductDetails(productId uint) (*models.ProductDetails, error) {
	return pds.productDetailsRep.GetProductDetailsByProductId(productId)
}

func (pds *ProductDetailsService) CreateProductDetails(productDetails *models.ProductDetails) error {
	return pds.productDetailsRep.CreateProductDetails(productDetails)
}

func (pds *ProductDetailsService) UpdateProductDetails(productDetails *models.ProductDetails) error {
	return pds.productDetailsRep.UpdateProductDetails(productDetails)
}

func (pds *ProductDetailsService) DeleteProductDetails(productId uint) error {
	return pds.productDetailsRep.DeleteProductDetails(productId)
}
