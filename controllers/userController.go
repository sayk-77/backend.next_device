package controllers

import (
	"github.com/gofiber/fiber/v2"
	"next_device/backend/models"
	"next_device/backend/service"
	"next_device/backend/tools"
	"strconv"
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

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)

	user, err := c.userService.GetUserById(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userResponse := models.UserProfileResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Addresses: user.Addresses,
		Orders:    user.Orders,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"user": userResponse})
}

func (c *UserController) SaveUserInfo(ctx *fiber.Ctx) error {
	var user *models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID := ctx.Locals("userID")
	if userID == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User ID not found in context"})
	}

	user.ID = userID.(uint)

	result, err := c.userService.SaveUserInfo(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *UserController) ChangePassword(ctx *fiber.Ctx) error {
	type userPassword struct {
		ID          uint   `json:"id"`
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	var req *userPassword
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID := ctx.Locals("userID")
	if userID == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User ID not found in context"})
	}

	req.ID = userID.(uint)

	result, err := c.userService.ChangePassword(req.ID, req.NewPassword, req.OldPassword)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *UserController) AddNewAddress(ctx *fiber.Ctx) error {
	var address *models.Address
	if err := ctx.BodyParser(&address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID := ctx.Locals("userID")
	if userID == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User ID not found in context"})
	}

	result, err := c.userService.AddNewAddress(userID.(uint), address)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (c *UserController) DeleteAddress(ctx *fiber.Ctx) error {
	addressId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID := ctx.Locals("userID")
	if userID == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User ID not found in context"})
	}

	result, err := c.userService.DeleteAddress(uint(addressId), userID.(uint))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
