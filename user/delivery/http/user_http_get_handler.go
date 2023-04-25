package http

import (
	_dto "microservice/shared/dto"
	_mapper "microservice/shared/pkg/mapper"

	"github.com/gofiber/fiber/v2"
)

func (u *UserHttpHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := u.UUsecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to get user", res))
	return err
}

func (u *UserHttpHandler) Fetch(c *fiber.Ctx) error {
	pagination := &_dto.Pagination{}

	if err := c.QueryParser(pagination); err != nil {
		return err
	}

	res, err := u.UUsecase.Fetch(c.Context(), *pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to get users", res))
	return err
}
