package routes

import (
	user_controller "restaurant_os/internal/api/user/controller"
	"restaurant_os/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(api fiber.Router) {

	protected := api.Group("", middleware.RequireAuth())
	users := protected.Group("/users")

	userHandler := user_controller.NewUserController()

	// User management routes
	users.Get("/", middleware.RequireRole("SUPER_ADMIN", "MANAGER"), userHandler.GetUsers)
	users.Post("/", middleware.RequireRole("SUPER_ADMIN", "MANAGER"), userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUser)
	users.Put("/:id", middleware.RequireRole("SUPER_ADMIN", "MANAGER"), userHandler.UpdateUser)
	users.Delete("/:id", middleware.RequireRole("SUPER_ADMIN", "MANAGER"), userHandler.DeleteUser)
	users.Get("/profile", userHandler.GetProfile)
	users.Put("/profile", userHandler.UpdateProfile)
}
