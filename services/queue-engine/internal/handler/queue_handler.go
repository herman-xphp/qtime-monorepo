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

type TakeTicketRequest struct {
	MerchantID string `json:"merchant_id"`
}

func (h *QueueHandler) TakeTicket(c *fiber.Ctx) error {
	var req TakeTicketRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.MerchantID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "merchant_id is required",
		})
	}

	// Call service
	number, err := h.svc.GenerateTicket(c.Context(), req.MerchantID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return JSON Response
	return c.JSON(fiber.Map{
		"status":        "success",
		"ticket_number": number,
		"merchant_id":   req.MerchantID,
	})
}
