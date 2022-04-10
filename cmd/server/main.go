package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	_config "microservice/shared/config"
	_mysql "microservice/shared/pkg/database/mysql"
	_redis "microservice/shared/pkg/database/redis"

	// HEALTH
	_healthHttpDelivery "microservice/health/delivery/http"

	// USER
	_userHttpDelivery "microservice/user/delivery/http"
	_mysqlUserRepository "microservice/user/repository/mysql"
	_userUsecase "microservice/user/usecase"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := _config.InitConfig(os.Getenv("ENV")); err != nil {
		panic(err)
	}

	if err := _mysql.Init(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")); err != nil {
		panic(err)
	}

	if err := _redis.Init(); err != nil {
		panic(err)
	}
}

func main() {
	mysqlConn := _mysql.MySQLManager()
	rdbConn := _redis.RedisManager()

	// Start a new fiber app
	app := fiber.New(fiber.Config{
		// Prefork:      true, // Need research
		WriteTimeout: 1 * time.Second, // Timeout after 1s
		ReadTimeout:  1 * time.Second, // Timeout after 1s
		IdleTimeout:  1 * time.Second, // Timeout after 1s
	})

	app.Use(cors.New()) // For CORS

	v1Grp := app.Group("/v1")

	// Health Group
	_healthHttpDelivery.NewHealthHandler(v1Grp)

	// User Group
	ur := _mysqlUserRepository.NewMysqlUserRepository(mysqlConn, rdbConn)
	uu := _userUsecase.NewUserUsecase(ur)
	_userHttpDelivery.NewUserHandler(v1Grp, uu)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = ":3000"
	}

	app.Listen(app_port)
}
