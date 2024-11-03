package controllers

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/models"
	"next_device/backend/service"
	"next_device/backend/tools"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := c.userService.Register(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := tools.CreateToken(user.ID, user.Role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Token: token,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered successfully", "user": userResponse})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	user, err := c.userService.Login(req.Email, req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := tools.CreateToken(user.ID, user.Role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := models.UserResponse{
		ID:    user.ID,
		Token: token,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful", "user": userResponse})
}
