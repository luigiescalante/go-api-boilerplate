package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func LimiterMiddleware() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        1000,
		Expiration: 60 * time.Second,
	})
}
