package http

import (
	"os"

	"github.com/gofiber/fiber/v2"

	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
	_helper "microservice/shared/pkg/helper"
	_mapper "microservice/shared/pkg/mapper"
	_middleware "microservice/shared/pkg/middleware"
	_validator "microservice/shared/pkg/validator"
)

// UserHandler represent the httphandler for server user
type UserHandler struct {
	UUsecase _domain.UserUsecase
}

// hltGrp.Get("", timeout.New(handler.Check, 5*time.Second)) // DON'T USE TIMEOUT (RACE CONDITION)
// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(router fiber.Router, us _domain.UserUsecase) {
	baseMiddleware := _middleware.New(_dto.Config{
		Filter: func(c *fiber.Ctx) bool {

			// if X-Secret-Pass header matches then we skip the jwt validation
			return c.Get("X-Secret-Pass") == os.Getenv("X_SECRET_PASS")

		},
	})
	handler := &UserHandler{
		UUsecase: us,
	}
	usrGrp := router.Group("/user", baseMiddleware)
	{
		usrGrp.Post("", handler.Store)
		usrGrp.Get("/:id", handler.GetByID)
		usrGrp.Get("", handler.Fetch)
		usrGrp.Put("", handler.Update)
		usrGrp.Delete("/:id", handler.Delete)
	}
	authGrp := router.Group("/auth")
	{
		authGrp.Post("/login", handler.Login)
		authGrp.Post("/refresh", handler.Refresh, baseMiddleware)
		authGrp.Post("/logout", handler.Logout, baseMiddleware)
		authGrp.Post("/reset_password", handler.ResetPassword, baseMiddleware)
	}
}

func (u *UserHandler) Store(c *fiber.Ctx) error {
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

func (u *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := u.UUsecase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to get user", res))
	return err
}

func (u *UserHandler) Fetch(c *fiber.Ctx) error {
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

func (u *UserHandler) Update(c *fiber.Ctx) error {
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

func (u *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.UUsecase.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(_mapper.BaseResponse("failed", err.Error(), nil))
	}

	err = c.Status(fiber.StatusOK).JSON(_mapper.BaseResponse("success", "success to update user", map[string]string{"user_id": id}))
	return err
}

func (u *UserHandler) Login(c *fiber.Ctx) error {
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

func (u *UserHandler) Refresh(c *fiber.Ctx) error {
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

func (u *UserHandler) Logout(c *fiber.Ctx) error {
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

func (u *UserHandler) ResetPassword(c *fiber.Ctx) error {
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
