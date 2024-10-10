package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type ProductService struct {
	productRepo  *repository.ProductRepository
	imageService *ProductImageService
}

func NewProductService(productRepo *repository.ProductRepository, imageService *ProductImageService) *ProductService {
	return &ProductService{productRepo: productRepo, imageService: imageService}
}

func (ps *ProductService) CreateProduct(product *models.Products) error {
	return ps.productRepo.CreateProduct(product)
}

func (ps *ProductService) GetAllProducts() ([]*models.Products, error) {
	products, err := ps.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProductByID(id uint) (*models.Products, error) {
	product, err := ps.productRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) UpdateProduct(product *models.Products) error {
	return ps.productRepo.UpdateProduct(product)
}

func (ps *ProductService) DeleteProduct(id uint) error {
	return ps.productRepo.DeleteProduct(id)
}

func (ps *ProductService) GetProductsByCategory(category string, limit, offset int) ([]*models.ProductWithMainImage, error) {
	products, err := ps.productRepo.GetProductsByCategoryPaged(category, limit, offset)
	if err != nil {
		return nil, err
	}

	var productsWithImages []*models.ProductWithMainImage
	for _, product := range products {
		mainImage, err := ps.imageService.GetMainImage(product.ID)
		if err != nil {
			return nil, err
		}
		productsWithImages = append(productsWithImages, &models.ProductWithMainImage{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Image:       mainImage.ImageURL,
			Price:       product.Price,
		})
	}

	return productsWithImages, nil
}

func (ps *ProductService) GetDiscountedProducts(limit, offset int) ([]*models.ProductWithMainImage, error) {
	products, err := ps.productRepo.GetDiscountedProductsPaged(limit, offset)
	if err != nil {
		return nil, err
	}

	var productsWithImages []*models.ProductWithMainImage
	for _, product := range products {
		mainImage, err := ps.imageService.GetMainImage(product.ID)
		if err != nil {
			return nil, err
		}
		productsWithImages = append(productsWithImages, &models.ProductWithMainImage{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Image:       mainImage.ImageURL,
			Price:       product.Price,
		})
	}

	return productsWithImages, nil
}

func (ps *ProductService) GetNewProducts(limit, offset int) ([]*models.ProductWithMainImage, error) {
	products, err := ps.productRepo.GetNewProductsPaged(limit, offset)
	if err != nil {
		return nil, err
	}

	var productsWithImages []*models.ProductWithMainImage
	for _, product := range products {
		mainImage, err := ps.imageService.GetMainImage(product.ID)
		if err != nil {
			return nil, err
		}
		productsWithImages = append(productsWithImages, &models.ProductWithMainImage{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Image:       mainImage.ImageURL,
			Price:       product.Price,
		})
	}

	return productsWithImages, nil
}
