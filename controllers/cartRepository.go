package controllers

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/models"
	"next_device/backend/service"
	"strconv"
)

type CartController struct {
	cartService *service.CartService
}

func NewCartController(cartService *service.CartService) *CartController {
	return &CartController{cartService}
}

func (c *CartController) AddItem(ctx *fiber.Ctx) error {
	var item models.CartItem
	if err := ctx.BodyParser(&item); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID := ctx.Locals("userID").(uint)

	if err := c.cartService.AddItem(userID, item.ProductId, item.Quantity); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not add item to cart"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item added to cart"})
}

func (c *CartController) RemoveItem(ctx *fiber.Ctx) error {
	productID, err := strconv.ParseUint(ctx.Params("productId"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	userID := ctx.Locals("userID").(uint)

	if err := c.cartService.RemoveItem(userID, uint(productID)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not remove item from cart"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item removed from cart"})
}

func (c *CartController) GetCartItems(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)

	items, err := c.cartService.GetCartItems(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not get cart items"})
	}

	return ctx.Status(fiber.StatusOK).JSON(items)
}
