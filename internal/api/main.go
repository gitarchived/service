package api

import (
	"log"
	"time"

	"github.com/gitarchived/service/internal/api/routes"
	"github.com/gitarchived/service/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Init() {
	app := fiber.New(fiber.Config{
		AppName: "GitArchived API",
	})

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 20 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{
				"message": "Too many requests",
				"status":  429,
			})
		},
	}))

	db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "GitArchived API",
			"status":  200,
		})
	})

	app.Post("/create", func(c *fiber.Ctx) error {
		return routes.Create(c, db)
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		return routes.Search(c, db)
	})

	app.Listen(":8080")
}
