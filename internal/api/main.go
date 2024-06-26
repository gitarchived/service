package api

import (
	"log"
	"time"

	"github.com/gitarchived/service/internal/api/middlewares"
	"github.com/gitarchived/service/internal/api/routes"
	"github.com/gitarchived/service/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Init() {
	app := fiber.New(fiber.Config{
		AppName: "GitArchived API",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal Server Error",
				"status":  500,
			})
		},
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

	app.Use(middlewares.Headers)

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

	app.Get("/:host/:owner/:name", func(c *fiber.Ctx) error {
		return routes.Get(c, db)
	})

	app.Get("/:host/:owner", func(c *fiber.Ctx) error {
		return routes.Owner(c, db)
	})

	// ! 404 handler needs to be at the end
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
			"status":  404,
		})
	})

	err = app.Listen(":8080")

	if err != nil {
		panic(err) // Panic because the application should not continue if the server fails to start
	}
}
