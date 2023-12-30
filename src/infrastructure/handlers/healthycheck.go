package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.api-boilerplate/services"
)

// HealthyCheck
//
//	@HealthyCheck	godoc
//	@Description	Get the healthy check status and ping services
//	@Tags			Healthy
//	@Produce		json
//	@Router 		/v1/healthy [get]
func HealthyCheck(c *fiber.Ctx) error {
	return WrapperSuccess(c, services.GetHealthy(), "")
}
