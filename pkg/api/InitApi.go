package api

import (
	"config"
	"strconv"

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

	host := config.GetConfig().Webserver.Host
	port := strconv.Itoa(config.GetConfig().Webserver.Port)

	err := app.Listen(host + ":" + port)
	if err != nil {
		panic(err)
	}

}
