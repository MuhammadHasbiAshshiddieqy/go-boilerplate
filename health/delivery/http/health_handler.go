package http

import (
	"github.com/gofiber/fiber/v2"
)

// HealthHandler represent the httphandler for server health
type HealthHandler struct{}

// hltGrp.Get("", timeout.New(handler.Check, 5*time.Second)) // DON'T USE TIMEOUT (RACE CONDITION)
// NewHealthHandler will initialize the health/ resources endpoint
func NewHealthHandler(router fiber.Router) {
	handler := &HealthHandler{}
	hltGrp := router.Group("/health")
	{
		hltGrp.Get("", handler.Check)
	}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "good condition"})
	return nil
}
