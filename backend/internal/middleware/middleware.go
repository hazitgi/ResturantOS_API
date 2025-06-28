package middleware

import (
	"restaurant_os/internal/dto"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"os"
)

func RequireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			errMsg := "Authorization header is missing or invalid"
			return c.Status(fiber.StatusUnauthorized).JSON(dto.APIResponse{
				Success: false,
				Message: "Unauthorized",
				Error:   &errMsg,
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			errMsg := "Token is missing"
			return c.Status(fiber.StatusUnauthorized).JSON(dto.APIResponse{
				Success: false,
				Message: "Unauthorized",
				Error:   &errMsg,
			})
		}

		accessSecret := os.Getenv("JWT_ACCESS_SECRET")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessSecret), nil // Use the actual secret from environment
		})
		if err != nil || !token.Valid {
			errMsg := "Invalid or expired token"
			return c.Status(fiber.StatusUnauthorized).JSON(dto.APIResponse{
				Success: false,
				Message: "Unauthorized",
				Error:   &errMsg,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			errMsg := "Invalid token claims"
			return c.Status(fiber.StatusUnauthorized).JSON(dto.APIResponse{
				Success: false,
				Message: "Unauthorized",
				Error:   &errMsg,
			})
		}

		// Set claims in context
		if userID, ok := claims["user_id"].(float64); ok {
			c.Locals("userID", uint(userID))
		}
		if email, ok := claims["email"].(string); ok {
			c.Locals("email", email)
		}
		if userType, ok := claims["user_type"].(string); ok {
			c.Locals("userType", userType)
		}
		if role, ok := claims["role"].(string); ok {
			c.Locals("role", role)
		}
		if restaurantID, ok := claims["restaurant_id"].(float64); ok {
			c.Locals("restaurantID", uint(restaurantID))
		}
		if branchID, ok := claims["branch_id"].(float64); ok {
			c.Locals("branchID", uint(branchID))
		}

		return c.Next()
	}
}

func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("userRole").(string)
		if !ok {
			errMsg := "User role not found in context"
			return c.Status(fiber.StatusUnauthorized).JSON(dto.APIResponse{
				Success: false,
				Message: "Unauthorized",
				Error:   &errMsg,
			})
		}

		for _, role := range allowedRoles {
			if userRole == role {
				return c.Next()
			}
		}

		errMsg := "You do not have permission to access this resource"
		return c.Status(fiber.StatusForbidden).JSON(dto.APIResponse{
			Success: false,
			Message: "Access denied",
			Error:   &errMsg,
		})
	}
}
