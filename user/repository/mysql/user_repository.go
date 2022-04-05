package repository

import (
	_domain "microservice/shared/domain"
	_dto "microservice/shared/dto"
	_helper "microservice/shared/pkg/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Orm *gorm.DB
}

// NewMysqlUserRepository will create an object that represent the userRepository interface
func NewMysqlUserRepository(Orm *gorm.DB) _domain.UserRepository {
	return &mysqlUserRepository{Orm}
}

func (u *mysqlUserRepository) Store(c *fiber.Ctx, us _domain.User) (_domain.User, error) {
	if err := u.Orm.Create(&us).Error; err != nil {
		return us, err
	}

	return us, nil
}

func (u *mysqlUserRepository) GetByID(c *fiber.Ctx, id string) (_domain.User, error) {
	us := _domain.User{}
	if err := u.Orm.First(&us, "id = ?", id).Error; err != nil {
		return us, err
	}

	return us, nil
}

func (u *mysqlUserRepository) Fetch(c *fiber.Ctx, pagination *_dto.Pagination) ([]*_domain.User, error) {
	var usrs []*_domain.User
	if err := u.Orm.Scopes(_helper.Paginate(usrs, pagination, u.Orm)).Find(&usrs); err != nil {
		return usrs, nil
	}

	return usrs, nil
}

func (u *mysqlUserRepository) Update(c *fiber.Ctx, us _domain.User) (_domain.User, error) {
	if err := u.Orm.Updates(&us).Error; err != nil {
		return us, err
	}

	return us, nil
}

func (u *mysqlUserRepository) Delete(c *fiber.Ctx, id string) error {
	if err := u.Orm.Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
