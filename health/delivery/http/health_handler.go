package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

// HealthHandler represent the httphandler for server health
type HealthHandler struct{}

// NewHealthHandler will initialize the health/ resources endpoint
func NewHealthHandler(router fiber.Router) {
	handler := &HealthHandler{}
	hltGrp := router.Group("/health")
	{
		hltGrp.Get("", timeout.New(handler.Check, 5*time.Second))
	}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "good condition"})
	return nil
}
