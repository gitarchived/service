package routes

import (
	"github.com/gitarchived/service/internal/db"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Owner(c *fiber.Ctx, db *db.DB) error {
	host := c.Params("host")
	owner := c.Params("owner")

	data, err := db.RepositoriesByOwner(host, owner)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"status":  404,
				"message": "Not Found",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"data":    data,
	})
}
