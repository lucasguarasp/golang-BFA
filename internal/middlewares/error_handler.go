package middlewares

import "github.com/gofiber/fiber/v3"

func ErrorHandler() fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}
}