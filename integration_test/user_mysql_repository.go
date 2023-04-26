package integrationtest

import (
	"context"
	_config "microservice/shared/config"
	_domain "microservice/shared/domain"

	// DB
	_mysql "microservice/shared/pkg/database/mysql"
	_userMysqlRepository "microservice/user/repository/mysql"

	"os"
	"testing"

	"github.com/joho/godotenv"
	// "github.com/stretchr/testify/assert"
)

func TestUserMysqlRepository_Store(t *testing.T) {
	var err error

	//To load our environmental variables.
	if err = godotenv.Load(); err != nil {
		panic(err)
	}

	if err = _config.InitConfig(os.Getenv("ENV")); err != nil {
		panic(err)
	}

	conf := _config.GetConfig()

	mysqlMaster := _mysql.New(
		conf.MysqlMicroMaster.Role,
		conf.MysqlMicroMaster.Port,
		conf.MysqlMicroMaster.Host,
		conf.MysqlMicroMaster.DB,
		os.Getenv("MYSQL_MICRO_MASTER_USERNAME"),
		os.Getenv("MYSQL_MICRO_MASTER_PASSWORD"),
		conf.MysqlMicroMaster.Name,
	)

	if err = mysqlMaster.Init(); err != nil {
		panic(err)
	}

	mysqlConn := _mysql.GetConnection(conf.MysqlMicroMaster.Name)
	umsqlr := _userMysqlRepository.NewUserMysqlRepository(mysqlConn)
	user := _domain.User{
		Name:     "Hasbi",
		Password: "password_test",
		Email:    "hasbi@gmail.com",
		RoleID:   "1",
	}
	user, err = umsqlr.Store(context.Background(), user)
}
