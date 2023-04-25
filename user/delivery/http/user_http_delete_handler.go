package http

import (
	_mapper "microservice/shared/pkg/mapper"

	"github.com/gofiber/fiber/v2"
)

func (u *UserHttpHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.UUsecase.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to update user", map[string]string{"user_id": id}))
	return err
}
