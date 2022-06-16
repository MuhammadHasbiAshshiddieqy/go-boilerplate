package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"microservice/cmd/server/register"
	_config "microservice/shared/config"
	_mysql "microservice/shared/pkg/database/mysql"
	_redis "microservice/shared/pkg/database/redis"
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
