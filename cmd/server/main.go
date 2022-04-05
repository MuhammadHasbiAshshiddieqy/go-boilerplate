package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	_config "microservice/shared/config"
	_mysql "microservice/shared/pkg/database/mysql"

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
		log.Println("no env detected")
	}
}

func main() {
	_config.InitConfig(os.Getenv("ENV"))

	_mysql.Init(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))

	// Start a new fiber app
	app := fiber.New()

	app.Use(cors.New()) // For CORS

	// Health Group
	_healthHttpDelivery.NewHealthHandler(app)

	// User Group
	ur := _mysqlUserRepository.NewMysqlUserRepository(_mysql.DbManager())
	uu := _userUsecase.NewUserUsecase(ur)
	_userHttpDelivery.NewUserHandler(app, uu)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = ":3000"
	}

	app.Listen(app_port)
}
