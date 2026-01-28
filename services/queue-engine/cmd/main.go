package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/config"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Fiber
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// 3. Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "queue-engine",
			"env":     cfg.App.Env,
		})
	})

	// 4. Start Server
	log.Fatal(app.Listen(cfg.HTTP.Port))
}
