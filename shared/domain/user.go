package domain

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"

	_dto "microservice/shared/dto"
)

type (
	User struct {
		Base
		Name     string `gorm:"column:name;type:varchar(64)"`
		Password string `gorm:"column:password;type:varchar(64)"`
	}

	UserUsecase interface {
		Fetch(c *fiber.Ctx, pagination _dto.Pagination) (_dto.Pagination, error)
		GetByID(c *fiber.Ctx, id string) (_dto.UserResponse, error)
		Update(c *fiber.Ctx, ureq _dto.UserRequestUpdate) (_dto.UserResponse, error)
		Store(c *fiber.Ctx, ureq _dto.UserRequestCreate) (_dto.UserResponse, error)
		Delete(c *fiber.Ctx, id string) error
		Login(c *fiber.Ctx, ureq _dto.UserRequestLogin) (_dto.UserResponseToken, error)
		Refresh(c *fiber.Ctx, ureq _dto.UserRequestRefresh) (_dto.UserResponseToken, error)
		Logout(c *fiber.Ctx, metadata *_dto.AccessDetails) error
		ResetPassword(c *fiber.Ctx, metadata *_dto.AccessDetails, ureq _dto.UserRequestPasswordUpdate) error
	}

	UserRepository interface {
		Fetch(c *fiber.Ctx, pagination *_dto.Pagination) ([]*User, error)
		GetByID(c *fiber.Ctx, id string) (User, error)
		GetByCondition(c *fiber.Ctx, u User) (User, error)
		Update(c *fiber.Ctx, u User) (User, error)
		Store(c *fiber.Ctx, u User) (User, error)
		Delete(c *fiber.Ctx, id string) error
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID, err = gonanoid.New()
	if err != nil {
		err = errors.New("failed to generate nano ID")
	}

	return
}
