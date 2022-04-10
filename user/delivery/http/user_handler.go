package http

import (
	"github.com/gofiber/fiber/v2"

	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
)

// UserHandler represent the httphandler for server health
type UserHandler struct {
	UUsecase _domain.UserUsecase
}

// hltGrp.Get("", timeout.New(handler.Check, 5*time.Second)) // DON'T USE TIMEOUT (RACE CONDITION)
// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(router fiber.Router, us _domain.UserUsecase) {
	handler := &UserHandler{
		UUsecase: us,
	}
	usrGrp := router.Group("/user")
	{
		usrGrp.Post("", handler.Store)
		usrGrp.Get("/:id", handler.GetByID)
		usrGrp.Get("", handler.Fetch)
		usrGrp.Put("", handler.Update)
		usrGrp.Delete("/:id", handler.Delete)
	}
}

func (u *UserHandler) Store(c *fiber.Ctx) error {
	payload := _dto.UserRequestCreate{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	res, err := u.UUsecase.Store(c, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to create user", "data": res})
	return nil
}

func (u *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := u.UUsecase.GetByID(c, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to get user", "data": res})
	return nil
}

func (u *UserHandler) Fetch(c *fiber.Ctx) error {
	pagination := &_dto.Pagination{}

	if err := c.QueryParser(pagination); err != nil {
		return err
	}

	res, err := u.UUsecase.Fetch(c, *pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to get users", "data": res})
	return nil
}

func (u *UserHandler) Update(c *fiber.Ctx) error {
	payload := _dto.UserRequestUpdate{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	res, err := u.UUsecase.Update(c, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to update user", "data": res})
	return nil
}

func (u *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := u.UUsecase.Delete(c, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to delete user", "user_id": id})
	return nil
}
