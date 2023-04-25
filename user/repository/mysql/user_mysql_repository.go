package mysqlrepository

import (
	_domain "microservice/shared/domain"

	"gorm.io/gorm"
)

type userMysqlRepository struct {
	Orm *gorm.DB
}

// NewUserMysqlRepository will create an object that represent the userRepository interface
func NewUserMysqlRepository(Orm *gorm.DB) _domain.UserMysqlRepository {
	return &userMysqlRepository{Orm: Orm}
}
