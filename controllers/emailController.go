package controllers

//
//import (
//	"github.com/gofiber/fiber/v2"
//	"next_device/backend/models"
//	"next_device/backend/service"
//)
//
//type EmailController struct {
//	emailService *service.EmailService
//}
//
//func NewEmailController(app *fiber.App, emailService *service.EmailService) *EmailController {
//	emailController := &EmailController{
//		emailService: emailService,
//	}
//	app.Post("/send-mail-order-status", emailController.SendMailOrderStatus)
//	app.Post("/send-mail-review-status", emailController.SendMailReviewStatus)
//	return emailController
//}
//
//func (ec *EmailController) SendMailOrderStatus(c *fiber.Ctx) error {
//	var email *models.Email
//	if err := c.BodyParser(&email); err != nil {
//		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
//	}
//
//	if err := ec.emailService.SendMail(email); err != nil {
//		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
//	}
//
//	return nil
//}
