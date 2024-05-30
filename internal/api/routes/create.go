package routes

import (
	"strings"

	"github.com/gitarchived/service/internal/db"
	"github.com/gitarchived/service/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateRequestBody struct {
	Url string `json:"url" validate:"required,http_url"`
}

func Create(c *fiber.Ctx, db *db.DB) error {
	body := new(CreateRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	// Split the URL into host, owner and name
	split := strings.Split(strings.TrimPrefix(strings.TrimSuffix(body.Url, "/"), "https://"), "/")

	if len(split) != 3 {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	host := split[0]
	owner := split[1]
	name := split[2]

	if !db.IsValidHost(host) {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	if !util.IsOk(body.Url) {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	if db.RepositoryExists(host, owner, name) {
		return c.Status(409).JSON(fiber.Map{
			"status":  409,
			"message": "Conflict",
		})
	}

	data, err := db.CreateRepository(host, owner, name)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  201,
		"message": "OK",
		"data":    data,
	})
}
