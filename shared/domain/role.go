package domain

import (
	"context"
	"errors"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"

	_dto "microservice/shared/dto"
)

type (
	Role struct {
		Base
		Name string `gorm:"column:name;type:varchar(64)"`
	}

	RoleUsecase interface {
		Fetch(c context.Context, pagination _dto.Pagination) (_dto.Pagination, error)
		GetByID(c context.Context, id string) (_dto.RoleResponse, error)
		Update(c context.Context, rreq _dto.RoleRequestUpdate) (_dto.RoleResponse, error)
		Store(c context.Context, rreq _dto.RoleRequestCreate) (_dto.RoleResponse, error)
		Delete(c context.Context, id string) error
	}

	RoleRepository interface {
		Fetch(c context.Context, pagination *_dto.Pagination) ([]*Role, error)
		GetByID(c context.Context, id string) (Role, error)
		Update(c context.Context, r Role) (Role, error)
		Store(c context.Context, r Role) (Role, error)
		Delete(c context.Context, id string) error
	}
)

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID, err = gonanoid.New()
	if err != nil {
		err = errors.New("failed to generate nano ID")
	}

	return
}
