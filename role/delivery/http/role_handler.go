package http

// import (
// 	"github.com/gofiber/fiber/v2"

// 	_domain "microservice/shared/domain"
// 	_dto "microservice/shared/dto"
// 	_helper "microservice/shared/pkg/helper"
// )

// // RoleHandler represent the httphandler for server role
// type RoleHandler struct {
// 	RUsecase _domain.RoleUsecase
// }

// // hltGrp.Get("", timeout.New(handler.Check, 5*time.Second)) // DON'T USE TIMEOUT (RACE CONDITION)
// // NewRoleHandler will initialize the user/ resources endpoint
// func NewRoleHandler(router fiber.Router, us _domain.RoleUsecase) {
// 	handler := &RoleHandler{
// 		RUsecase: us,
// 	}
// 	roleGrp := router.Group("/role")
// 	{
// 		roleGrp.Post("", handler.Store)
// 		roleGrp.Get("/:id", handler.GetByID)
// 		roleGrp.Get("", handler.Fetch)
// 		roleGrp.Put("", handler.Update)
// 		roleGrp.Delete("/:id", handler.Delete)
// 	}
// }

// func (u *UserHandler) Store(c *fiber.Ctx) error {
// 	payload := _dto.UserRequestCreate{}
// 	if err := c.BodyParser(&payload); err != nil {
// 		return err
// 	}

// 	res, err := u.UUsecase.Store(c.Context(), payload)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
// 	}

// 	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to create user", "data": res})
// 	return err
// }

// func (u *UserHandler) GetByID(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	res, err := u.UUsecase.GetByID(c.Context(), id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
// 	}

// 	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to get user", "data": res})
// 	return err
// }

// func (u *UserHandler) Fetch(c *fiber.Ctx) error {
// 	pagination := &_dto.Pagination{}

// 	if err := c.QueryParser(pagination); err != nil {
// 		return err
// 	}

// 	res, err := u.UUsecase.Fetch(c.Context(), *pagination)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
// 	}

// 	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to get users", "data": res})
// 	return err
// }

// func (u *UserHandler) Update(c *fiber.Ctx) error {
// 	payload := _dto.UserRequestUpdate{}
// 	if err := c.BodyParser(&payload); err != nil {
// 		return err
// 	}

// 	res, err := u.UUsecase.Update(c.Context(), payload)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
// 	}

// 	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to update user", "data": res})
// 	return err
// }

// func (u *UserHandler) Delete(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	err := u.UUsecase.Delete(c.Context(), id)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "failed", "message": err.Error()})
// 	}

// 	err = c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "success to delete user", "user_id": id})
// 	return err
// }
