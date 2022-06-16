package register

import (
	"github.com/gofiber/fiber/v2"

	// DB
	_mysql "microservice/shared/pkg/database/mysql"
	_redis "microservice/shared/pkg/database/redis"

	// HEALTH
	_healthHttpDelivery "microservice/health/delivery/http"

	// USER
	_userHttpDelivery "microservice/user/delivery/http"
	_mysqlUserRepository "microservice/user/repository/mysql"
	_userUsecase "microservice/user/usecase"
)

func InitV1(v1Grp fiber.Router) {
	mysqlConn := _mysql.MySQLManager()
	rdbConn := _redis.RedisManager()

	// Health Group
	_healthHttpDelivery.NewHealthHandler(v1Grp)

	// User Group
	ur := _mysqlUserRepository.NewMysqlUserRepository(mysqlConn, rdbConn)
	uu := _userUsecase.NewUserUsecase(ur)
	_userHttpDelivery.NewUserHandler(v1Grp, uu)
}
