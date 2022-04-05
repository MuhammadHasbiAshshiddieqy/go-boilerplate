package repository

import (
	_domain "microservice/shared/domain"

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

func (u *mysqlUserRepository) Update(c *fiber.Ctx, us _domain.User) (_domain.User, error) {
	if err := u.Orm.Updates(&us).Error; err != nil {
		return us, err
	}

	return us, nil
}
