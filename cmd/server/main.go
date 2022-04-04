package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	_config "microservice/shared/config"
	// _mysql "microservice/shared/pkg/database/mysql"

	_healthHttpDelivery "microservice/health/delivery/http"
	_healthUsecase "microservice/health/usecase"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	_config.InitConfig(os.Getenv("ENV"))

	// err := _mysql.Init(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	// if err != nil {
	// 	panic(err)
	// }

	// Start a new fiber app
	app := fiber.New()

	app.Use(cors.New()) // For CORS

	// Health Group
	au := _healthUsecase.NewHealthUsecase()
	_healthHttpDelivery.NewHealthHandler(app, au)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = ":3000"
	}

	app.Listen(app_port)
}
