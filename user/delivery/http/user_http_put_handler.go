package http

import (
	_dto "microservice/shared/dto"
	_mapper "microservice/shared/pkg/mapper"
	_validator "microservice/shared/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

// @Summary Update
// @Description Update User Data
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.UserRequestUpdate true "request body"
// @Success 200 {object} dto.BaseResponse
// @Security ApiKeyAuth
// @Router /users [put]
func (u *UserHttpHandler) Update(c *fiber.Ctx) error {
	payload := _dto.UserRequestUpdate{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	validateErr := _validator.ValidateStruct(payload)
	if validateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(_mapper.BaseResponse("failed", "failed to validate request", validateErr))

	}

	res, err := u.UUsecase.Update(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to update user", res))
	return err
}
