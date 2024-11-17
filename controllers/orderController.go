package controllers

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/service"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{orderService}
}

func (c *OrderController) GetOrderById(ctx *fiber.Ctx) error {
	orderId, err := strconv.Atoi(ctx.Params("id"))

	order, err := c.orderService.FindOrderById(uint(orderId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(order)
}

func (c *OrderController) GetAllOrders(ctx *fiber.Ctx) error {
	orders, err := c.orderService.GetAllOrders()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(orders)
}
