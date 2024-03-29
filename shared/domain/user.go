package domain

import (
	"context"
	"errors"

	_helper "microservice/shared/pkg/helper"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"

	_dto "microservice/shared/dto"
)

type (
	User struct {
		Base
		Name     string `gorm:"column:name;type:varchar(64)"`
		Password string `gorm:"column:password;type:varchar(64)"`
		Email    string `gorm:"column:email;type:varchar(64)"`
		RoleID   string `gorm:"column:role_id;type:varchar(64)"`
		Role     Role   `gorm:"->;foreignKey:RoleID;references:ID"`
	}

	UserUsecase interface {
		Fetch(c context.Context, pagination _dto.Pagination) (_dto.Pagination, error)
		GetByID(c context.Context, id string) (_dto.UserResponse, error)
		Update(c context.Context, ureq _dto.UserRequestUpdate) (_dto.UserResponse, error)
		Store(c context.Context, ureq _dto.UserRequestCreate) (_dto.UserResponse, error)
		Delete(c context.Context, id string) error
		Login(c context.Context, ureq _dto.UserRequestLogin) (_dto.UserResponseToken, error)
		Refresh(c context.Context, ureq _dto.UserRequestRefresh) (_dto.UserResponseToken, error)
		Logout(c context.Context, metadata *_dto.AccessDetails) error
		ResetPassword(c context.Context, metadata *_dto.AccessDetails, ureq _dto.UserRequestPasswordUpdate) error
	}

	UserMysqlRepository interface {
		Fetch(c context.Context, pagination *_dto.Pagination) ([]*User, error)
		GetByID(c context.Context, id string) (User, error)
		GetByCondition(c context.Context, u User) (User, error)
		Update(c context.Context, u User) (User, error)
		Store(c context.Context, u User) (User, error)
		Delete(c context.Context, id string) error
	}

	UserRedisRepository interface {
		GetByID(c context.Context, id string) (User, error)
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID, err = gonanoid.New()
	if err != nil {
		err = errors.New("failed to generate nano ID")
	}
	u.Password, err = _helper.HashPassword(u.Password)
	if err != nil {
		err = errors.New("failed to hash password")
	}

	return
}
