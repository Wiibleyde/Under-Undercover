package api

import (
	"game"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func playTurnDescriptionApi(c *fiber.Ctx) error {
	gameUuid := strings.Clone(c.Query("gameUuid"))
	if gameUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de partie fourni",
			},
			GameData: ResponseGameData{},
		}
		return c.JSON(response)
	}
	playerUuid := c.Query("playerUuid")
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: ResponseGameData{},
		}
		return c.JSON(response)
	}
	wordGiven := c.Query("wordGiven")
	if wordGiven == "" {
		response := Response{
			Message: Message{
				Error: "Aucun mot donné",
			},
			GameData: ResponseGameData{},
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
		}
		return c.JSON(response)
	}

	descErr := gameSelected.PlayTurnDesc(player, wordGiven)
	if descErr != nil {
		response := Response{
			Message: Message{
				Error: descErr.Error(),
			},
			GameData: ResponseGameData{},
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

	responseGameData := ResponseGameData{
		GameUuid:   gameSelected.Uuid,
		Started:    gameSelected.Started,
		Ended:      gameSelected.Ended,
		PlayerTurn: gameSelected.PlayerTurn,
		Players:    players,
		PlaysDesc:  playsDesc,
		PlaysVote:  []PlaysVoteDataHidden{},
		GameState:  gameSelected.GameState,
	}

	response := Response{
		Message: Message{
			Information: "Tour joué",
		},
		GameData: responseGameData,
		UserWord: gameSelected.GetPlayerWord(player),
	}
	return c.JSON(response)
}

func playTurnDiscussionApi(c *fiber.Ctx) error {
	gameUuid := c.Query("gameUuid")
	if gameUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de partie fourni",
			},
			GameData: ResponseGameData{},
		}
		return c.JSON(response)
	}
	playerUuid := c.Query("playerUuid")
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: ResponseGameData{},
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
		}
		return c.JSON(response)
	}

	descErr := gameSelected.PlayTurnDiscuss(player)
	if descErr != nil {
		response := Response{
			Message: Message{
				Error: descErr.Error(),
			},
			GameData: ResponseGameData{},
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

	responseGameData := ResponseGameData{
		GameUuid:   gameSelected.Uuid,
		Started:    gameSelected.Started,
		Ended:      gameSelected.Ended,
		PlayerTurn: gameSelected.PlayerTurn,
		Players:    players,
		PlaysDesc:  playsDesc,
		PlaysVote:  []PlaysVoteDataHidden{},
		GameState:  gameSelected.GameState,
	}

	response := Response{
		Message: Message{
			Information: "Tour joué",
		},
		GameData: responseGameData,
		UserWord: gameSelected.GetPlayerWord(player),
	}
	return c.JSON(response)
}

func playTurnVoteApi(c *fiber.Ctx) error {
	gameUuid := c.Query("gameUuid")
	if gameUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de partie fourni",
			},
			GameData: ResponseGameData{},
		}
		return c.JSON(response)
	}
	playerUuid := c.Query("playerUuid")
	if playerUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de joueur fourni",
			},
			GameData: ResponseGameData{},
		}
		return c.JSON(response)
	}
	voteUuid := c.Query("voteUuid")
	if voteUuid == "" {
		response := Response{
			Message: Message{
				Error: "Aucun identifiant de vote fourni",
			},
			GameData: ResponseGameData{},
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
		}
		return c.JSON(response)
	}

	playerVoted, err := gameSelected.GetPlayer(voteUuid)
	if err != nil {
		response := Response{
			Message: Message{
				Error: err.Error(),
			},
			GameData: ResponseGameData{},
		}
		return c.JSON(response)
	}

	voteErr := gameSelected.PlayTurnElim(player, playerVoted)
	if voteErr != nil {
		response := Response{
			Message: Message{
				Error: voteErr.Error(),
			},
			GameData: ResponseGameData{},
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
			Information: "Tour joué",
		},
		GameData: responseGameData,
		UserWord: gameSelected.GetPlayerWord(player),
	}
	return c.JSON(response)
}
