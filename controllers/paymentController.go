package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"log"
	"next_device/backend/models"
	"next_device/backend/service"
	"os"
	"strconv"
)

type PaymentController struct {
	paymentService *service.PaymentService
	orderService   *service.OrderService
}

func NewPaymentController(paymentService *service.PaymentService, orderService *service.OrderService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
		orderService:   orderService,
	}
}

func (c *PaymentController) HandlePaymentIntent(ctx *fiber.Ctx) error {
	var req struct {
		OrderItems []models.OrderItem `json:"orderItems"`
		TotalPrice float64            `json:"totalPrice"`
		Address    int                `json:"address"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	userID := ctx.Locals("userID").(uint)

	order, err := c.orderService.CreateOrder(userID, req.OrderItems, req.TotalPrice, uint(req.Address))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}

	session, err := c.paymentService.CreateCheckoutSession(order.ID, req.TotalPrice)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create payment session"})
	}

	return ctx.JSON(fiber.Map{"id": session.ID})
}

func (c *PaymentController) HandlePaymentStatus(ctx *fiber.Ctx) error {
	var req struct {
		OrderId   string `json:"orderId"`
		SessionId string `json:"sessionId"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Error parsing body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	stripeSession, err := session.Get(req.SessionId, nil)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve session"})
	}

	orderId, err := strconv.Atoi(req.OrderId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid orderId"})
	}

	if stripeSession.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid {
		payment := models.Payment{
			OrderID:       uint(orderId),
			PaymentID:     req.SessionId,
			PaymentMethod: "credit_card",
			Amount:        float64(stripeSession.AmountTotal),
			PaymentStatus: "completed",
		}
		if err := c.paymentService.SavePayment(&payment); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save payment"})
		}
	} else {
		if err := c.orderService.DeleteOrder(uint(orderId)); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete order"})
		}
	}

	return ctx.JSON(fiber.Map{"message": "Payment status processed"})
}
