package api

import (
	"game"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func createGameApi(c *fiber.Ctx) error {
	gameSelected := game.Game{}
	gameSelected.CreateGame()

	responseGameData := ResponseGameData{
		GameUuid:   gameSelected.Uuid,
		Started:    gameSelected.Started,
		Ended:      gameSelected.Ended,
		PlayerTurn: gameSelected.PlayerTurn,
		Players:    []PlayerHidden{},
		PlaysDesc:  []PlaysDescDataHidden{},
		PlaysVote:  []PlaysVoteDataHidden{},
		GameState:  gameSelected.GameState,
	}

	response := Response{
		Message: Message{
			Information: "Partie créée",
		},
		GameData: responseGameData,
		UserWord: "",
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
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}
	playerUuid := strings.Clone(c.Query("playerUuid"))
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}
	playerPseudo := strings.Clone(c.Query("playerPseudo"))
	if playerPseudo == "" {
		response := Response{
			Message: Message{
				Error: "Aucun pseudo de joueur fourni",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	gameSelected, err := game.GetGame(gameUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
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
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	var players []PlayerHidden
	for _, player := range gameSelected.Players {
		players = append(players, PlayerHidden{
			Pseudo:     player.Pseudo,
			Eliminated: player.Eliminated,
			Connected:  player.Connected,
		})
	}

	responseGameData := ResponseGameData{
		GameUuid:   gameSelected.Uuid,
		Started:    gameSelected.Started,
		Ended:      gameSelected.Ended,
		PlayerTurn: gameSelected.PlayerTurn,
		Players:    players,
		PlaysDesc:  []PlaysDescDataHidden{},
		PlaysVote:  []PlaysVoteDataHidden{},
		GameState:  gameSelected.GameState,
	}

	response := Response{
		Message: Message{
			Information: "Joueur ajouté à la partie",
		},
		GameData: responseGameData,
		UserWord: "",
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
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}
	playerUuid := strings.Clone(c.Query("playerUuid"))
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	gameSelected, err := game.GetGame(gameUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	err = gameSelected.RemovePlayer(playerUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	var players []PlayerHidden
	for _, player := range gameSelected.Players {
		players = append(players, PlayerHidden{
			Pseudo:     player.Pseudo,
			Eliminated: player.Eliminated,
			Connected:  player.Connected,
		})
	}

	responseGameData := ResponseGameData{
		GameUuid:   gameSelected.Uuid,
		Started:    gameSelected.Started,
		Ended:      gameSelected.Ended,
		PlayerTurn: gameSelected.PlayerTurn,
		Players:    players,
		PlaysDesc:  []PlaysDescDataHidden{},
		PlaysVote:  []PlaysVoteDataHidden{},
		GameState:  gameSelected.GameState,
	}

	response := Response{
		Message: Message{
			Information: "Joueur retiré de la partie",
		},
		GameData: responseGameData,
		UserWord: "",
	}
	return c.JSON(response)
}

func startGameApi(c *fiber.Ctx) error {
	gameUuid := strings.Clone(c.Query("gameUuid"))
	if gameUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de partie fourni",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}
	playerUuid := strings.Clone(c.Query("playerUuid"))
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	gameSelected, err := game.GetGame(gameUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	if gameSelected.Started {
		response := Response{
			Message: Message{
				Error: "La partie a déjà commencé",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	player, err := gameSelected.GetPlayer(playerUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	gameSelected.InitGame(player)
	err = gameSelected.StartGame()
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	var players []PlayerHidden
	for _, player := range gameSelected.Players {
		players = append(players, PlayerHidden{
			Pseudo:     player.Pseudo,
			Eliminated: player.Eliminated,
			Connected:  player.Connected,
		})
	}

	responseGameData := ResponseGameData{
		GameUuid:   gameSelected.Uuid,
		Started:    gameSelected.Started,
		Ended:      gameSelected.Ended,
		PlayerTurn: gameSelected.PlayerTurn,
		Players:    players,
		PlaysDesc:  []PlaysDescDataHidden{},
		PlaysVote:  []PlaysVoteDataHidden{},
		GameState:  gameSelected.GameState,
	}

	response := Response{
		Message: Message{
			Information: "Partie démarrée",
		},
		GameData: responseGameData,
		UserWord: gameSelected.GetPlayerWord(player),
	}
	return c.JSON(response)
}

func getGameApi(c *fiber.Ctx) error {
	gameUuid := strings.Clone(c.Query("gameUuid"))
	if gameUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de partie fourni",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	playerUuid := strings.Clone(c.Query("playerUuid"))
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	gameSelected, err := game.GetGame(gameUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	player, err := gameSelected.GetPlayer(playerUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
			UserWord: "",
		}
		return c.JSON(response)
	}

	var players []PlayerHidden
	for _, player := range gameSelected.Players {
		players = append(players, PlayerHidden{
			Pseudo:     player.Pseudo,
			Eliminated: player.Eliminated,
			Connected:  player.Connected,
		})
	}

	var playsDesc []PlaysDescDataHidden
	for _, play := range gameSelected.PlaysDesc {
		playsDesc = append(playsDesc, PlaysDescDataHidden{
			PlayerUuid: play.Player.Uuid,
			Word:       play.WordGiven,
		})
	}

	var playsVote []PlaysVoteDataHidden
	for _, vote := range gameSelected.PlaysVote {
		playsVote = append(playsVote, PlaysVoteDataHidden{
			PlayerUuid: vote.Player.Uuid,
			VoteUuid:   vote.Vote.Uuid,
		})
	}

	responseGameData := ResponseGameData{
		GameUuid:   gameSelected.Uuid,
		Started:    gameSelected.Started,
		Ended:      gameSelected.Ended,
		PlayerTurn: gameSelected.PlayerTurn,
		Players:    players,
		PlaysDesc:  playsDesc,
		PlaysVote:  playsVote,
		GameState:  gameSelected.GameState,
	}

	response := Response{
		Message: Message{
			Information: "Partie récupérée",
		},
		GameData: responseGameData,
		UserWord: gameSelected.GetPlayerWord(player),
	}

	return c.JSON(response)
}
