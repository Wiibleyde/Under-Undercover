package api

import (
	"config"
	"logger"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// InitApi initializes the API
func InitApi() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error { return c.SendStatus(200) })
	app.Get("/api", func(c *fiber.Ctx) error { return c.SendString("Welcome to the API !") })

	app.Get("/api/createGame", createGameApi)
	app.Get("/api/joinGame", joinGameApi)
	app.Get("/api/leaveGame", leaveGameApi)
	app.Get("/api/startGame", startGameApi)

	app.Get("/api/getGame", getGameApi)

	app.Post("/api/playTurn/description", playTurnDescriptionApi)
	app.Post("/api/playTurn/discussion", playTurnDiscussionApi)
	app.Post("/api/playTurn/vote", playTurnVoteApi)

	host := config.GetConfig().Webserver.Host
	port := strconv.Itoa(config.GetConfig().Webserver.Port)

	err := app.Listen(host + ":" + port)
	if err != nil {
		logger.ErrorLogger.Panicln("Error while runing the API: ", err)
	}

}
