package usecase

import (
	"github.com/gofiber/fiber/v2"

	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
)

type healthUsecase struct{}

// NewHealthUsecase will create new an healthUsecase object representation of domain.HealthUsecase interface
func NewHealthUsecase() _domain.HealthUsecase {
	return &healthUsecase{}
}

func (h *healthUsecase) Check(c *fiber.Ctx) (_dto.BaseResponse, error) {
	return _dto.BaseResponse{Message: "Success"}, nil
}
