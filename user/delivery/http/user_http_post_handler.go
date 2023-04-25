package http

import (
	_dto "microservice/shared/dto"
	_helper "microservice/shared/pkg/helper"
	_mapper "microservice/shared/pkg/mapper"
	_validator "microservice/shared/pkg/validator"

	"github.com/gofiber/fiber/v2"
)

func (u *UserHttpHandler) Store(c *fiber.Ctx) error {
	payload := _dto.UserRequestCreate{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	validateErr := _validator.ValidateStruct(payload)
	if validateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(_mapper.BaseResponse("failed", "failed to validate request", validateErr))

	}

	res, err := u.UUsecase.Store(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to create user", res))
	return err
}

func (u *UserHttpHandler) Login(c *fiber.Ctx) error {
	payload := _dto.UserRequestLogin{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	validateErr := _validator.ValidateStruct(payload)
	if validateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(_mapper.BaseResponse("failed", "failed to validate request", validateErr))

	}

	res, err := u.UUsecase.Login(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "login success", res))
	return err
}

func (u *UserHttpHandler) Refresh(c *fiber.Ctx) error {
	payload := _dto.UserRequestRefresh{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	validateErr := _validator.ValidateStruct(payload)
	if validateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(_mapper.BaseResponse("failed", "failed to validate request", validateErr))

	}

	res, err := u.UUsecase.Refresh(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "refresh success", res))
	return err
}

func (u *UserHttpHandler) Logout(c *fiber.Ctx) error {
	metadata, err := _helper.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}
	err = u.UUsecase.Logout(c.Context(), metadata)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "logout success", nil))
	return err
}

func (u *UserHttpHandler) ResetPassword(c *fiber.Ctx) error {
	payload := _dto.UserRequestPasswordUpdate{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	validateErr := _validator.ValidateStruct(payload)
	if validateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(_mapper.BaseResponse("failed", "failed to validate request", validateErr))

	}

	metadata, err := _helper.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = u.UUsecase.ResetPassword(c.Context(), metadata, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "reset password success", nil))
	return err
}
