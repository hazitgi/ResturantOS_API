package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/hazitgi/restaurant_os/internal/models"

	"github.com/lpernett/godotenv"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	fmt.Println("Starting")
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	// Connect to the database
	err = models.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}


	app := fiber.New(fiber.Config{
		ServerHeader:  "Restaurant OS",
		AppName:       "Restaurant OS v0.1",
		CaseSensitive: true,
	})

	app.Use(healthcheck.New())
	app.Get("/monitor", monitor.New(monitor.Config{
		Title:       "Restaurant OS Monitor",
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

	log.Fatal(app.Listen(":5000"))

}
