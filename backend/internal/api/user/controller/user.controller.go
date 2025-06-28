package user_controller

import (
	"github.com/gofiber/fiber/v2"
)

type userController struct{}

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get all users"})
}

func (u *userController) CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User created"})
}

func (u *userController) GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get user by ID", "id": c.Params("id")})
}

func (u *userController) UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update user", "id": c.Params("id")})
}

func (u *userController) DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Delete user", "id": c.Params("id")})
}

func (u *userController) GetProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get current user's profile"})
}

func (u *userController) UpdateProfile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Update current user's profile"})
}
