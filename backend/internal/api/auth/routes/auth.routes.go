package routes

import (
	"github.com/gofiber/fiber/v2"
	auth_controller "restaurant_os/internal/api/auth/controller"
)

func RegisterAuthRoutes(api fiber.Router) {

	auth := api.Group("/auth")

	authController := auth_controller.NewAuthController()

	auth.Post("/login", authController.LoginHandler)
	auth.Post("/refresh", authController.RefreshTokenHandler)

}
