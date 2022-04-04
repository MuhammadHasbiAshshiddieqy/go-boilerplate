package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"

	_domain "microservice/shared/domain"
)

// HealthHandler represent the httphandler for server health
type HealthHandler struct {
	HUsecase _domain.HealthUsecase
}

// NewHealthHandler will initialize the health/ resources endpoint
func NewHealthHandler(router fiber.Router, he _domain.HealthUsecase) {
	handler := &HealthHandler{
		HUsecase: he,
	}
	hltGrp := router.Group("/health")
	{
		hltGrp.Get("", timeout.New(handler.Check, 5*time.Second))
	}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	ok, err := h.HUsecase.Check(c)
	if err != nil {
		c.JSON(err.Error())
		return err
	}
	c.JSON(ok)
	return nil
}
