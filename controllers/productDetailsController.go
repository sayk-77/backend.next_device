package controllers

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/models"
	"next_device/backend/service"
	"strconv"
)

type ProductDetailsController struct {
	productDetailsService *service.ProductDetailsService
}

func NewProductDetailsController(productDetailsService *service.ProductDetailsService) *ProductDetailsController {
	return &ProductDetailsController{productDetailsService: productDetailsService}
}

func (pdc *ProductDetailsController) GetProductDetails(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("productId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product id"})
	}

	productDetails, err := pdc.productDetailsService.GetProductDetails(uint(productId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Product not found"})
	}

	return c.JSON(productDetails)
}

func (pdc *ProductDetailsController) CreateProductDetails(c *fiber.Ctx) error {
	var productDetails *models.ProductDetails
	if err := c.BodyParser(&productDetails); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product details data",
		})
	}

	err := pdc.productDetailsService.CreateProductDetails(productDetails)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"result": "success"})
}

func (pdc *ProductDetailsController) UpdateProductDetails(c *fiber.Ctx) error {
	var productDetails *models.ProductDetails
	if err := c.BodyParser(&productDetails); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product details data",
		})
	}

	err := pdc.productDetailsService.UpdateProductDetails(productDetails)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"result": "success"})
}

func (pdc *ProductDetailsController) DeleteProductDetails(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("productId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product id",
		})
	}
	err = pdc.productDetailsService.DeleteProductDetails(uint(productId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"result": "success"})
}
