package http

import (
	"github.com/gofiber/fiber/v2"

	_mapper "microservice/shared/pkg/mapper"
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
	err := c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "server in good condition", nil))
	return err
}
