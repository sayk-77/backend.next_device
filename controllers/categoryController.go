package controllers

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/models"
	"next_device/backend/service"
	"strconv"
)

type CategoryController struct {
	categoryService *service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (cc *CategoryController) GetAllCategories(c *fiber.Ctx) error {
	category, err := cc.categoryService.GetAllCategory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

func (cc *CategoryController) GetCategoryByIdOrName(c *fiber.Ctx) error {
	param := c.Params("param")

	id, err := strconv.Atoi(param)
	if err == nil {
		brand, err := cc.categoryService.GetCategoryById(uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Category not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(brand)
	}

	category, err := cc.categoryService.GetCategoryByName(param)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

func (cc *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category data",
		})
	}

	if err := cc.categoryService.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
	})
}

func (cc *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category data",
		})
	}
	if err := cc.categoryService.UpdateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category updated successfully",
	})
}

func (cc *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category id",
		})
	}

	if err := cc.categoryService.DeleteCategoryById(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}

func (cc *CategoryController) GetAllCategoryAndCount(c *fiber.Ctx) error {
	category, err := cc.categoryService.GetAllCategoryAndCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve category",
		})
	}
	return c.Status(fiber.StatusOK).JSON(category)
}
