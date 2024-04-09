package api

import (
	"github.com/gofiber/fiber/v2"
)

// InitApi initializes the API
func InitApi() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		panic(err)
	}

}
