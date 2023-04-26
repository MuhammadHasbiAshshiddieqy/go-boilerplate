package http

import (
	_dto "microservice/shared/dto"
	_mapper "microservice/shared/pkg/mapper"

	"github.com/gofiber/fiber/v2"
)

// @Summary GetByID
// @Description Get User By ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (u *UserHttpHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := u.UUsecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to get user", res))
	return err
}

// @Summary Fetch
// @Description Get Users List
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query string true "data per page"
// @Param page query string true "page number"
// @Param sort query string true "asc/desc"
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Router /users [get]
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
