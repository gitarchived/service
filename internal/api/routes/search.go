package routes

import (
	"strconv"

	"github.com/gitarchived/service/internal/db"
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx, db *db.DB) error {
	query := c.Query("q")
	index, err := strconv.Atoi(c.Query("index"))

	if err != nil || index < 0 {
		index = 1
	}

	if query == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	data, con, err := db.SearchRepositories(query, index)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   200,
		"message":  "OK",
		"data":     data,
		"index":    index,
		"continue": con,
	})
}
