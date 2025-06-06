package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"next_device/backend/models"
	"next_device/backend/service"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ReviewController struct {
	reviewService *service.ReviewService
	pushService   *service.NotificationService
	orderService  *service.OrderService
	userService   *service.UserService
	emailService  *service.EmailService
}

func NewReviewController(reviewService *service.ReviewService, pushService *service.NotificationService, orderService *service.OrderService, userService *service.UserService, emailService *service.EmailService) *ReviewController {
	return &ReviewController{reviewService: reviewService, pushService: pushService, orderService: orderService, userService: userService, emailService: emailService}
}

func generateUniqueName() string {
	u := uuid.New().String()

	return strings.ReplaceAll(u, "-", "")
}

func (c *ReviewController) CreateReview(ctx *fiber.Ctx) error {
	var review models.Review
	if err := ctx.BodyParser(&review); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID := ctx.Locals("userID")
	if userID == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User ID not found in context"})
	}
	review.UserID = userID.(uint)

	if err := c.reviewService.CreateReview(&review); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var reviewImages []models.ReviewImage
	savePath := "./images_review"

	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		log.Println("Error creating directory:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось создать папку для загрузки файлов"})
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Ошибка при обработке изображений"})
	}

	files := form.File["images"]

	for _, file := range files {
		uniqueName := fmt.Sprintf("review_%s%s", generateUniqueName(), filepath.Ext(file.Filename))
		filePath := filepath.Join(savePath, uniqueName)

		if err := ctx.SaveFile(file, filePath); err != nil {
			log.Printf("Error saving file %s: %v", file.Filename, err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ошибка при сохранении файла"})
		}

		reviewImage := models.ReviewImage{
			ReviewId: review.ID,
			ImageUrl: uniqueName,
		}
		reviewImages = append(reviewImages, reviewImage)
	}

	if err := c.reviewService.CreateReviewImages(reviewImages); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Отзыв успешно создан"})
}

func (c *ReviewController) GetReviewById(ctx *fiber.Ctx) error {
	reviewId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	review, err := c.reviewService.GetReviewById(uint(reviewId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(review)
}

func (c *ReviewController) DeleteReview(ctx *fiber.Ctx) error {
	reviewId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	review, err := c.reviewService.GetReviewById(uint(reviewId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	userID := int(review.UserID)

	user, err := c.userService.GetUserById(uint(userID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	result := c.reviewService.DeleteReview(uint(reviewId))

	err = c.emailService.SendReviewStatusEmail(user.Email, user.FirstName, "был отклонен, так как он содержил недопустимый контент")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	notif := models.Notification{
		Title: "Отзыв",
		Body:  "Ваш отзыв был отклонен, так как он содержал недопустимый контент",
		Icon:  "/logo_not.webp",
		Link:  "http://localhost:3000/profile",
	}

	if err := c.pushService.SendNotificationToUser(userID, notif); err != nil {
		log.Printf("Не удалось отправить уведомление пользователю %d: %v", userID, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось отправить уведомление",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": result})
}

func (c *ReviewController) PublishReview(ctx *fiber.Ctx) error {
	reviewId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result := c.reviewService.PublishReview(uint(reviewId))

	review, err := c.reviewService.GetReviewById(uint(reviewId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	userID := int(review.UserID)

	user, err := c.userService.GetUserById(uint(userID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	err = c.emailService.SendReviewStatusEmail(user.Email, user.FirstName, "был опубликован. Теперь его видят все пользователи")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	notif := models.Notification{
		Title: "Отзыв",
		Body:  "Ваш отзыв был опубликован",
		Icon:  "/logo_not.webp",
		Link:  "http://localhost:3000/profile",
	}

	if err := c.pushService.SendNotificationToUser(userID, notif); err != nil {
		log.Printf("Не удалось отправить уведомление пользователю %d: %v", userID, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось отправить уведомление",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": result})

}

func (c *ReviewController) GetReviewForProduct(ctx *fiber.Ctx) error {
	productId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	reviews, err := c.reviewService.GetReviewForProduct(uint(productId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(reviews)
}

func (c *ReviewController) GetAllReviews(ctx *fiber.Ctx) error {
	review, err := c.reviewService.GetAllReviews()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return ctx.JSON(review)
}

func (c *ReviewController) ChangeStatus(ctx *fiber.Ctx) error {
	var req struct {
		OrderId int    `json:"orderId"`
		Status  string `json:"status"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.reviewService.ChangeStatus(uint(req.OrderId), req.Status); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	order, err := c.orderService.FindOrderById(uint(req.OrderId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	userID := int(order.UserID)

	user, err := c.userService.GetUserById(uint(userID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	err = c.emailService.SendOrderStatusEmail(user.Email, user.FirstName, req.Status, req.OrderId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	notif := models.Notification{
		Title: "Заказ",
		Body:  "Статус вашего заказа изменен",
		Icon:  "/logo_not.webp",
		Link:  "http://localhost:3000/profile",
	}

	if err := c.pushService.SendNotificationToUser(userID, notif); err != nil {
		log.Printf("Не удалось отправить уведомление пользователю %d: %v", userID, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось отправить уведомление",
		})
	}

	return ctx.JSON(fiber.Map{"status": "success"})
}
