package controllers

import (
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

	c.service.AddSubscription(sub)

	return ctx.JSON(fiber.Map{
		"message": "Подписка успешно сохранена",
	})
}

func (c *NotificationController) GetPublicKey(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"publicKey": c.service.GetPublicKey(),
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
