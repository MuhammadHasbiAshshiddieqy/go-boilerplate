package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"microservice/cmd/server/register"
	_config "microservice/shared/config"
	_domain "microservice/shared/domain"
	_mysql "microservice/shared/pkg/database/mysql"
	_redis "microservice/shared/pkg/database/redis"
)

func init() {
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

	if err = _mysql.GetConnection(conf.MysqlMicroMaster.Name).AutoMigrate(
		&_domain.User{},
		&_domain.Role{},
	); err != nil {
		panic(err)
	}

	redisMaster := _redis.New(
		conf.RedisMaster.Role,
		conf.RedisMaster.Port,
		conf.RedisMaster.DB,
		conf.RedisMaster.Host,
		os.Getenv("REDIS_MASTER_PASSWORD"),
		conf.RedisMaster.Name,
	)

	if err = redisMaster.Init(); err != nil {
		panic(err)
	}

	fmt.Println("MySQL connection : ", _mysql.GetConnectionList())
	fmt.Println("Redis connection : ", _redis.GetConnectionList())
}

func main() {
	// Start a new fiber app
	app := fiber.New(fiber.Config{
		// Prefork:      true, // Need research
		WriteTimeout: 1 * time.Second, // Timeout after 1s
		ReadTimeout:  1 * time.Second, // Timeout after 1s
		IdleTimeout:  1 * time.Second, // Timeout after 1s
	})

	app.Use(cors.New()) // For CORS

	v1Grp := app.Group("/v1")
	register.InitV1(v1Grp)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = ":3000"
	}

	err := app.Listen(app_port)
	if err != nil {
		panic(err)
	}
}
