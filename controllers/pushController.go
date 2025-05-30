package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"next_device/backend/models"
	"next_device/backend/service"
)

type NotificationController struct {
	service *service.NotificationService
}

func NewNotificationController(service *service.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

func (c *NotificationController) Subscribe(ctx *fiber.Ctx) error {
	var sub models.Subscription
	if err := ctx.BodyParser(&sub); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	userID := ctx.Locals("userID").(uint)

	c.service.AddSubscription(&sub, int(userID))

	return ctx.JSON(fiber.Map{
		"message": "Подписка успешно сохранена",
	})
}

func (c *NotificationController) SendNotification(ctx *fiber.Ctx) error {
	var notif models.Notification
	if err := ctx.BodyParser(&notif); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	if err := c.service.SendNotificationToAll(notif); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при отправке уведомлений",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Уведомления отправлены всем подписчикам",
	})
}

func (c *NotificationController) SendNotificationToUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("userId", "-1")

	var userIDInt int
	if _, err := fmt.Sscanf(userID, "%d", &userIDInt); err != nil || userIDInt <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный идентификатор пользователя",
		})
	}

	var notif models.Notification
	if err := ctx.BodyParser(&notif); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	if err := c.service.SendNotificationToUser(userIDInt, notif); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при отправке уведомления пользователю",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Уведомление отправлено пользователю",
	})
}
