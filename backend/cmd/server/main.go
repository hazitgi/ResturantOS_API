package main

import (
	"fmt"
	"log"

	"restaurant_os/internal/config"
	"restaurant_os/internal/dto"
	"restaurant_os/internal/models"
	"restaurant_os/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	fmt.Println("Starting")
	// Load environment variables from .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	// Connect to the database
	err = models.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	app := fiber.New(fiber.Config{
		ServerHeader:  "Restaurant OS",
		AppName:       "Restaurant OS v0.1",
		CaseSensitive: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			errMsg := err.Error()
			return c.Status(code).JSON(dto.APIResponse{
				Success: false,
				Message: "Internal Server Error",
				Error:   &errMsg,
			})
		},
	})
	// Global middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(healthcheck.New())
	app.Get("/monitor", monitor.New(monitor.Config{
		Title: "Restaurant OS Monitor",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		result := map[string]interface{}{
			"message": "Hello World",
			"version": "v1.0.0",
			"status":  200,
			"data": map[string]interface{}{
				"first_name": "hazi",
				"last_name":  "tgi",
				"email":      "hazi.tgi@gmail.com",
				"age":        25,
				"address": map[string]interface{}{
					"city":    "Baku",
					"country": "Azerbaijan",
				},
			},
		}

		ctx.Set("Content-Type", "application/graphql-response+json")
		return ctx.JSON(result)
	})

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":5000"))

}
