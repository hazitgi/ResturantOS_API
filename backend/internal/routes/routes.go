package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "restaurant_os/internal/api/auth/routes"
	user "restaurant_os/internal/api/user/routes"
)

func RegisterRoutes(app *fiber.App) {
	// Initialize all route groups
	api := app.Group("/api/v1")

	auth.RegisterAuthRoutes(api)
	user.RegisterUserRoutes(api)

}
