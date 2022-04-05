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
		// 	Fetch(ctx context.Context, cursor string, num int64) ([]User, string, error)
		GetByID(c *fiber.Ctx, id string) (_dto.UserResponse, error)
		Update(c *fiber.Ctx, u _dto.UserRequestUpdate) (_dto.UserResponse, error)
		Store(c *fiber.Ctx, u _dto.UserRequestCreate) (_dto.UserResponse, error)
		// 	Delete(ctx context.Context, id int64) error
	}

	UserRepository interface {
		// 	Fetch(ctx context.Context, cursor string, num int64) ([]User, string, error)
		GetByID(c *fiber.Ctx, id string) (User, error)
		Update(c *fiber.Ctx, u User) (User, error)
		Store(c *fiber.Ctx, u User) (User, error)
		// 	Delete(ctx context.Context, id int64) error
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID, err = gonanoid.New()
	if err != nil {
		err = errors.New("failed to generate nano ID")
	}

	return
}
