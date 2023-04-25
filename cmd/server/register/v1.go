package register

import (
	_config "microservice/shared/config"

	"github.com/gofiber/fiber/v2"

	// DB
	_mysql "microservice/shared/pkg/database/mysql"
	_redis "microservice/shared/pkg/database/redis"

	// HEALTH
	_healthHttpDelivery "microservice/health/delivery/http"

	// USER
	_userHttpDelivery "microservice/user/delivery/http"
	_userMysqlRepository "microservice/user/repository/mysql"
	_userRedisRepository "microservice/user/repository/redis"
	_userUsecase "microservice/user/usecase"
)

func InitV1(v1Grp fiber.Router) {
	conf := _config.GetConfig()
	mysqlConn := _mysql.GetConnection(conf.MysqlMicroMaster.Name)
	rdbConn := _redis.GetConnection(conf.RedisMaster.Name)

	// Health Group
	_healthHttpDelivery.NewHealthHandler(v1Grp)

	// User Group
	umsqlr := _userMysqlRepository.NewUserMysqlRepository(mysqlConn)
	urdsr := _userRedisRepository.NewUserRedisRepository(rdbConn)
	uu := _userUsecase.NewUserUsecase(umsqlr, urdsr)
	_userHttpDelivery.NewUserHttpHandler(v1Grp, uu)
}
