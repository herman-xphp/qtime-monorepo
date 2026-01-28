package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/service"
)

type QueueHandler struct {
	svc *service.QueueService
}

func NewQueueHandler(svc *service.QueueService) *QueueHandler {
	return &QueueHandler{svc: svc}
}

func (h *QueueHandler) TakeTicket(c *fiber.Ctx) error {
	// Simple validation: Merchant ID required
	merchantID := c.Query("merchant_id")
	if merchantID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "merchant_id is required",
		})
	}

	// Call service
	number, err := h.svc.GenerateTicket(c.Context(), merchantID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return JSON Response
	return c.JSON(fiber.Map{
		"status":        "success",
		"ticket_number": number,
		"merchant_id":   merchantID,
	})
}
