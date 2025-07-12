package user_controller

import (
	"encoding/json"
	user_dto "restaurant_os/internal/api/user/dto"
	user_service "restaurant_os/internal/api/user/services"
	dto "restaurant_os/internal/dto"
	"restaurant_os/internal/models"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userController struct{}

var validate = validator.New()

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get all users"})
}

func (u *userController) CreateUser(c *fiber.Ctx) error {
	var req user_dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body", "details": err.Error()})
	}

	err := validate.Struct(&req)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string)
		for _, e := range errs {
			field := e.Field()
			msg, ok := user_dto.CreateUserValidationErrorMessages[field]
			if !ok {
				msg = "Invalid value"
			}
			validationErrors[strings.ToLower(field)] = msg
		}
		validationErrorsJSON, _ := json.Marshal(validationErrors)
		validationErrorsStr := string(validationErrorsJSON)
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success: false,
			Message: "Validation failed",
			Error:   &validationErrorsStr,
		})
	}

	// Assuming user service is available to handle user creation
	user, err := user_service.CreateUser(&req)
	if err != nil {
		errMsg := err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(dto.APIResponse{
			Success: false,
			Message: "Failed to create user",
			Error:   &errMsg,
		})
	}

	userResponse := user_dto.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		UserType: string(user.UserType),
		Role: func(role *models.EmployeeRole) *string {
			if role == nil {
				return nil
			}
			str := string(*role)
			return &str
		}(user.Role),
		RestaurantID: user.RestaurantID,
		BranchID:     user.BranchID,
		IsActive:     user.IsActive,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(dto.APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    userResponse,
	})
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
