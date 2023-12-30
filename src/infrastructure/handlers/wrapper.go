package handlers

import "github.com/gofiber/fiber/v2"

func WrapperSuccess(c *fiber.Ctx, data interface{}, message string) error {
	c.Status(fiber.StatusOK)
	return c.JSON(responseFormant(data, message))
}

func WrapperCreated(c *fiber.Ctx, data interface{}, message string) error {
	c.Status(fiber.StatusCreated)
	return c.JSON(responseFormant(data, message))
}

func WrapperInternalServerError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusInternalServerError)
	return c.JSON(responseFormant(nil, err.Error()))
}

func WrapperNotFoundError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusNotFound)
	return c.JSON(responseFormant(nil, err.Error()))
}

func WrapperForbiddenError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusForbidden)
	return c.JSON(responseFormant(nil, err.Error()))
}

func WrapperBadRequestError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusBadRequest)
	return c.JSON(responseFormant(nil, err.Error()))
}

func responseFormant(data interface{}, message string) map[string]interface{} {
	response := make(map[string]interface{})
	if data != nil {
		response["data"] = data
	}
	if message != "" {
		response["message"] = message
	}
	return response
}
