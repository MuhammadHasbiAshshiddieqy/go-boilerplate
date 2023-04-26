package http

import (
	"os"

	"github.com/gofiber/fiber/v2"

	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
	_middleware "microservice/shared/pkg/middleware"
)

// UserHandler represent the httphandler for server user
type UserHttpHandler struct {
	UUsecase _domain.UserUsecase
}

// hltGrp.Get("", timeout.New(handler.Check, 5*time.Second)) // DON'T USE TIMEOUT (RACE CONDITION)
// NewUserHttpHandler will initialize the user/ resources endpoint
func NewUserHttpHandler(router fiber.Router, us _domain.UserUsecase) {
	baseMiddleware := _middleware.New(_dto.Config{
		Filter: func(c *fiber.Ctx) bool {

			// if X-Secret-Pass header matches then we skip the jwt validation
			return c.Get("X-Secret-Pass") == os.Getenv("X_SECRET_PASS")

		},
	})
	handler := &UserHttpHandler{
		UUsecase: us,
	}
	usrGrp := router.Group("/users", baseMiddleware)
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
