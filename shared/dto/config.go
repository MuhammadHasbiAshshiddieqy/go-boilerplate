package dto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	// when returned true, our middleware is skipped
	Filter func(c *fiber.Ctx) bool

	// function to run when there is error decoding jwt
	Unauthorized fiber.Handler

	// function to decode our jwt token
	Decode func(c *fiber.Ctx) (*jwt.MapClaims, error)
}
