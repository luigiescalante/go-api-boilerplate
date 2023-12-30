package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"go.api-boilerplate/config"
	"go.api-boilerplate/infrastructure/middlewares"
)

func Handler() *fiber.App {
	app := fiber.New()
	swaggerDoc()
	app.Use(middlewares.LimiterMiddleware())
	apiV1 := app.Group("v1")
	apiV1.Get("/healthy", HealthyCheck)             //healthy check
	apiV1.Get("/swagger/*", swagger.HandlerDefault) // api doc
	// login
	apiV1.Post("/signup", SignUp)
	apiV1.Post("/login", Login)
	// JWT Middleware
	app.Use(middlewares.JwtAuthMiddleware())

	return app
}

func swaggerDoc() {
	cfg := config.GetConfig()
	docs.SwaggerInfo.Title = "Cicada Challenge"
	docs.SwaggerInfo.Description = "Cicada challenge for backend software engineer"
	docs.SwaggerInfo.BasePath = cfg.GetDomain()
	docs.SwaggerInfo.BasePath = "/"
}
