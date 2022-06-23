package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	_dto "microservice/shared/dto"
	_helper "microservice/shared/pkg/helper"
)

var ConfigDefault = _dto.Config{
	Filter:       nil,
	Decode:       nil,
	Unauthorized: nil,
}

func configDefault(config ..._dto.Config) _dto.Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values if not passed
	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	// this is the main jwt decode function of our middleware
	if cfg.Decode == nil {
		// Set default Decode function if not passed
		cfg.Decode = func(c *fiber.Ctx) (*jwt.MapClaims, error) {

			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return nil, errors.New("Authorization header is required")
			}

			// we parse our jwt token and check for validity against our secret
			token, err := _helper.VerifyToken(c)

			if err != nil {
				return nil, errors.New("Error parsing token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !(ok && token.Valid) {
				return nil, errors.New("Invalid token")
			}

			if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				return nil, errors.New("jwt is expired")
			}

			if roleId, ok := claims["role_id"]; ok && roleId.(string) != "1" {
				return nil, errors.New("unauthorized role")
			}

			return &claims, nil
		}
	}

	// Set default Unauthorized if not passed
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}

func New(config _dto.Config) fiber.Handler {

	// For setting default config
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Filter returns true
		if cfg.Filter != nil && cfg.Filter(c) {
			fmt.Println("Middleware was skipped")
			return c.Next()
		}
		fmt.Println("Middleware was run")

		claims, err := cfg.Decode(c)

		if err == nil {
			c.Locals("jwtClaims", *claims)
			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}
