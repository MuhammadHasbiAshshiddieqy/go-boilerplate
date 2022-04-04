package domain

import (
	"github.com/gofiber/fiber/v2"
)

type (
	HealthUsecase interface {
		Check(c *fiber.Ctx) error
	}
)
