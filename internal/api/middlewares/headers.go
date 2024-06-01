package middlewares

import "github.com/gofiber/fiber/v2"

func Headers(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	return c.Next()
}
