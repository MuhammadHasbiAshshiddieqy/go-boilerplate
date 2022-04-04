package domain

import (
	"github.com/gofiber/fiber/v2"

	_dto "microservice/shared/dto"
)

type (
	HealthUsecase interface {
		Check(c *fiber.Ctx) (_dto.BaseResponse, error)
	}
)
