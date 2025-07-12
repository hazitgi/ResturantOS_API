package controller

import (
	"fmt"
	auth_dto "restaurant_os/internal/api/auth/dto"
	"restaurant_os/internal/api/auth/helpers"
	auth_services "restaurant_os/internal/api/auth/services"
	user_dto "restaurant_os/internal/api/user/dto"
	dto "restaurant_os/internal/dto"
	"restaurant_os/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authController struct{}

func NewAuthController() *authController {
	return &authController{}
}

func (ac *authController) LoginHandler(c *fiber.Ctx) error {
	var loginRequest auth_dto.LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		errMsg := err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(dto.APIResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   &errMsg,
		})
	}

	fmt.Println("Login request received:", loginRequest)

	user, err := auth_services.GetUserByEmail(loginRequest.Email)

	if err != nil {
		errMsg := "Internal server error"
		if err.Error() == "user not found" {
			errMsg = "Invalid email or password"
			return c.Status(fiber.StatusUnauthorized).JSON(dto.APIResponse{
				Success: false,
				Message: errMsg,
				Error:   &errMsg,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.APIResponse{
			Success: false,
			Message: errMsg,
			Error:   &errMsg,
		})
	}

	if !helpers.CheckPasswordHash(loginRequest.Password, user.Password) {
		errMsg := "Invalid email or password"
		return c.Status(fiber.StatusUnauthorized).JSON(dto.APIResponse{
			Success: false,
			Message: errMsg,
			Error:   &errMsg,
		})
	}

	// Generate tokens
	accessToken := helpers.GenerateAccessToken(user)
	refreshToken := helpers.GenerateRefreshToken(user)

	if accessToken == "" || refreshToken == "" {
		errMsg := "Failed to generate tokens"
		return c.Status(fiber.StatusInternalServerError).JSON(dto.APIResponse{
			Success: false,
			Message: errMsg,
			Error:   &errMsg,
		})
	}

	userSummary := user_dto.UserSummary{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		UserType: string(user.UserType),
		Role: func(role *models.EmployeeRole) *string {
			if role == nil {
				return nil
			}
			str := string(*role)
			return &str
		}(user.Role),
	}

	response := &auth_dto.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		User:         userSummary,
		ExpiresAt:    time.Now(),
	}

	return c.JSON(dto.APIResponse{
		Success: true,
		Message: "Login successful",
		Data:    response,
	})
}

func (ac *authController) RefreshTokenHandler(c *fiber.Ctx) error {
	// Mock refresh logic
	response := &auth_dto.LoginResponse{
		Token:        "newToken",
		RefreshToken: "newRefreshToken",
		User:         user_dto.UserSummary{},
		ExpiresAt:    time.Now().Add(time.Hour),
	}

	return c.JSON(response)
}
