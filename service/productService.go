package service

import (
	"next_device/backend/models"
	"next_device/backend/repository"
)

type ProductService struct {
	productRepo  *repository.ProductRepository
	imageService *ProductImageService
	brandRepo    *repository.BrandRepository
	categoryRepo *repository.CategoryRepository
}

type SearchResults struct {
	Products   []*models.ProductWithMainImage `json:"products"`
	Brands     []*models.Brand                `json:"brands"`
	Categories []*models.Category             `json:"categories"`
}

func NewProductService(productRepo *repository.ProductRepository, imageService *ProductImageService, brandRepo *repository.BrandRepository, categoryRepo *repository.CategoryRepository) *ProductService {
	return &ProductService{productRepo: productRepo, imageService: imageService, brandRepo: brandRepo, categoryRepo: categoryRepo}
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

func (ps *ProductService) GetProductByName(name string) (*models.Products, error) {
	product, err := ps.productRepo.GetProductByName(name)
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
	products, title, err := ps.productRepo.GetProductsByCategoryPaged(category, limit, offset)
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
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			SearchName:    product.SearchName,
			CategoryTitle: title,
			DiscountPrice: product.DiscountPrice,
			Image:         mainImage.ImageURL,
			Price:         product.Price,
		})
	}

	return productsWithImages, nil
}

func (ps *ProductService) GetDiscountedProducts(limit, offset int, brand string) ([]*models.ProductWithMainImage, error) {
	products, err := ps.productRepo.GetDiscountedProductsPaged(limit, offset, brand)
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
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			SearchName:    product.SearchName,
			DiscountPrice: product.DiscountPrice,
			Image:         mainImage.ImageURL,
			Price:         product.Price,
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
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			SearchName:    product.SearchName,
			DiscountPrice: product.DiscountPrice,
			Image:         mainImage.ImageURL,
			Price:         product.Price,
		})
	}

	return productsWithImages, nil
}

func (ps *ProductService) GetProductsByBrandAndCategory(brandId int, categoryName string, limit, offset int) ([]*models.ProductWithMainImage, error) {
	product, err := ps.productRepo.GetProductsByBrandAndCategory(uint(brandId), categoryName, limit, offset)
	if err != nil {
		return nil, err
	}

	var productsWithImages []*models.ProductWithMainImage
	for _, product := range product {
		mainImage, err := ps.imageService.GetMainImage(product.ID)
		if err != nil {
			return nil, err
		}
		productsWithImages = append(productsWithImages, &models.ProductWithMainImage{
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			SearchName:    product.SearchName,
			DiscountPrice: product.DiscountPrice,
			Image:         mainImage.ImageURL,
			Price:         product.Price,
		})
	}
	return productsWithImages, nil
}

func (ps *ProductService) SearchAll(query string, limit, offset int) (*SearchResults, error) {
	results := SearchResults{}

	products, err := ps.productRepo.SearchProduct(query, limit, offset)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		mainImage, err := ps.imageService.GetMainImage(product.ID)
		if err != nil {
			return nil, err
		}
		results.Products = append(results.Products, &models.ProductWithMainImage{
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			SearchName:    product.SearchName,
			DiscountPrice: product.DiscountPrice,
			Image:         mainImage.ImageURL,
			Price:         product.Price,
		})
	}

	brands, err := ps.brandRepo.SearchBrand(query, limit, offset)
	if err != nil {
		return nil, err
	}
	results.Brands = brands

	categories, err := ps.categoryRepo.SearchCategory(query)
	if err != nil {
		return nil, err
	}
	results.Categories = categories

	return &results, nil
}

func (ps *ProductService) GetFilteredProducts(
	category string,
	priceFrom, priceTo *int,
	brands []string,
	screenFrom, screenTo *float64, // Параметры для экрана
	memories, ram, ratings, cameraQualities, os []string, // Добавлен os
	limit, offset int,
) ([]*models.ProductWithMainImage, error) {

	// Вызов метода репозитория с добавленным параметром os
	products, err := ps.productRepo.GetFilteredProducts(
		category, priceFrom, priceTo, brands, screenFrom, screenTo, memories, ram, ratings, cameraQualities, os, limit, offset,
	)
	if err != nil {
		return nil, err
	}

	// Собираем продукты с изображениями
	var productsWithImages []*models.ProductWithMainImage
	for _, product := range products {
		mainImage, err := ps.imageService.GetMainImage(product.ID)
		if err != nil {
			return nil, err
		}
		productsWithImages = append(productsWithImages, &models.ProductWithMainImage{
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			SearchName:    product.SearchName,
			DiscountPrice: product.DiscountPrice,
			Image:         mainImage.ImageURL,
			Price:         product.Price,
		})
	}

	return productsWithImages, nil
}
