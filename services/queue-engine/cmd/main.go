package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/config"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/handler"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/repository"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/service"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Connect to Redis
	rdb := config.NewRedisClient(cfg)
	defer rdb.Close()

	// 3. Wiring Dependency Injection (Clean Arch)
	// Repo -> Service -> Handler
	repo := repository.NewQueueRepository(rdb)
	svc := service.NewQueueService(repo)
	h := handler.NewQueueHandler(svc)

	// 4. Initialize Fiber
	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// 5. Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "queue-engine",
			"env":     cfg.App.Env,
			"redis":   "Connected",
		})
	})

	// API Endpoint for get queue
	app.Post("/queue/take", h.TakeTicket)

	// 6. Start Server
	log.Printf("🚀 Server starting on %s", cfg.HTTP.Port)
	log.Fatal(app.Listen(cfg.HTTP.Port))
}
