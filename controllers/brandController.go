package controllers

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/models"
	"next_device/backend/service"
	"strconv"
)

type BrandController struct {
	brandService *service.BrandService
}

func NewBrandController(brandService *service.BrandService) *BrandController {
	return &BrandController{brandService: brandService}
}

func (bc *BrandController) CreateBrand(c *fiber.Ctx) error {
	var brand models.Brand

	if err := c.BodyParser(&brand); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid brand data",
		})
	}

	if err := bc.brandService.CreateBrand(&brand); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create brand",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Brand created successfully",
	})
}

func (bc *BrandController) GetAllBrand(c *fiber.Ctx) error {

	limitQuery := c.Query("limit")
	var limit *int

	if limitQuery != "" {
		parsedLimit, err := strconv.Atoi(limitQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid limit parameter",
			})
		}
		limit = &parsedLimit
	}

	products, err := bc.brandService.GetAllBrand(limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve products brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

func (bc *BrandController) GetBrandByNameOrId(c *fiber.Ctx) error {
	param := c.Params("param")

	id, err := strconv.Atoi(param)
	if err == nil {
		brand, err := bc.brandService.GetBrandByID(uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Brand not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(brand)
	}

	brand, err := bc.brandService.GetBrandByName(param)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Brand not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(brand)
}

func (bc *BrandController) UpdateBrand(c *fiber.Ctx) error {
	var brand models.Brand

	if err := c.BodyParser(&brand); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid brand data",
		})
	}

	if err := bc.brandService.UpdateBrand(&brand); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brand updated successfully",
	})
}

func (bc *BrandController) DeleteBrand(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid brand ID",
		})
	}

	if err := bc.brandService.DeleteBrand(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete brand",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brand deleted successfully",
	})
}

func (bc *BrandController) GetCategoriesByBrand(c *fiber.Ctx) error {
	brandID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid brand ID",
		})
	}

	categories, err := bc.brandService.GetCategoriesByBrand(uint(brandID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve categories"})
	}

	return c.JSON(fiber.Map{
		"categories": categories,
	})
}
