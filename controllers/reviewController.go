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
}

func NewReviewController(reviewService *service.ReviewService) *ReviewController {
	return &ReviewController{reviewService: reviewService}
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

	result := c.reviewService.DeleteReview(uint(reviewId))

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": result})
}

func (c *ReviewController) PublishReview(ctx *fiber.Ctx) error {
	reviewId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result := c.reviewService.PublishReview(uint(reviewId))

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

	return ctx.JSON("success")
}
