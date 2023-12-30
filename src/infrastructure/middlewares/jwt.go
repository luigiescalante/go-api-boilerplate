package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.api-boilerplate/config"
)

type TokenClimbs struct {
	UserId uint64
}

func JwtAuthMiddleware() fiber.Handler {
	cfg := config.GetConfig()
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.GetJwtKey())},
	})
}

func GetTokenClaims(c *fiber.Ctx) *TokenClimbs {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)
	return &TokenClimbs{
		UserId: uint64(userId),
	}
}
