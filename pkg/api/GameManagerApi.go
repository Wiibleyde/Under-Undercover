package api

import (
	"game"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func createGameApi(c *fiber.Ctx) error {
	gameSelected := game.Game{}
	gameSelected.CreateGame()

	response := Response{
		Message: Message{
			Information: "Partie créée",
		},
		GameData: gameSelected,
	}
	return c.JSON(response)
}

func joinGameApi(c *fiber.Ctx) error {
	gameUuid := strings.Clone(c.Query("gameUuid"))
	if gameUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de partie fourni",
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}
	playerUuid := strings.Clone(c.Query("playerUuid"))
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}
	playerPseudo := strings.Clone(c.Query("playerPseudo"))
	if playerPseudo == "" {
		response := Response{
			Message: Message{
				Error: "Aucun pseudo de joueur fourni",
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	gameSelected, err := game.GetGame(gameUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	player := game.Player{
		Uuid:      playerUuid,
		Pseudo:    playerPseudo,
		Role:      game.NotSet,
		Connected: true,
	}

	err = gameSelected.AddPlayer(player)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	response := Response{
		Message: Message{
			Information: "Joueur ajouté à la partie",
		},
		GameData: gameSelected,
	}
	return c.JSON(response)
}

func leaveGameApi(c *fiber.Ctx) error {
	gameUuid := strings.Clone(c.Query("gameUuid"))
	if gameUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de partie fourni",
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}
	playerUuid := strings.Clone(c.Query("playerUuid"))
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	gameSelected, err := game.GetGame(gameUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	err = gameSelected.RemovePlayer(playerUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	response := Response{
		Message: Message{
			Information: "Joueur retiré de la partie",
		},
		GameData: gameSelected,
	}
	return c.JSON(response)
}

func startGameApi(c *fiber.Ctx) error {
	gameUuid := strings.Clone(c.Query("gameUuid"))

	gameSelected, err := game.GetGame(gameUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	err = gameSelected.StartGame()
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: game.Game{},
		}
		return c.JSON(response)
	}

	response := Response{
		Message: Message{
			Information: "Partie démarrée",
		},
		GameData: gameSelected,
	}
	return c.JSON(response)
}
