package main

import (
	"fmt"
	"game"
	"math/rand"
	"time"
	"words"
)

func main() {
	words.InitWords()

	player1 := game.Player{
		Uuid:   "1",
		Pseudo: "player1",
		// Position: -1,
		Role: game.NotSet,
	}
	player2 := game.Player{
		Uuid:   "2",
		Pseudo: "player2",
		// Position: -1,
		Role: game.NotSet,
	}
	player3 := game.Player{
		Uuid:   "3",
		Pseudo: "player3",
		// Position: -1,
		Role: game.NotSet,
	}
	player4 := game.Player{
		Uuid:   "4",
		Pseudo: "player4",
		// Position: -1,
		Role: game.NotSet,
	}

	game1 := game.Game{}
	game1.CreateGame()
	game1.AddPlayer(player1)
	game1.AddPlayer(player2)
	game1.AddPlayer(player3)
	game1.AddPlayer(player4)

	game1.InitGame(game1.Players[0])

	fmt.Println(game1)

	game1.StartGame()

	var turn int

	for {
		for game1.GameState.DescriptionPhase {
			playerTurn, _ := game1.GetNextPlayer()
			err2 := game1.PlayTurnDesc(playerTurn, "word2")
			if err2 != nil {
				// _ = a
				fmt.Println(err2.Error(), " ", fmt.Sprint(game1.GameState))
			} else {
				fmt.Println("Player", playerTurn.Pseudo, "gave a word", game1.PlayerTurn)
			}
		}

		err2 := game1.PlayTurnDiscuss(game1.Host)
		if err2 != nil {
			if err2.Error() == game.NotHost.Message {
				return
			}
			fmt.Println("HOST", err2, " ", fmt.Sprint(game1.GameState))
		}

		randomVote := rand.Intn(len(game1.Players))
		for game1.GameState.EliminationPhase {
			playerTurn, _ := game1.GetNextPlayer()
			err2 := game1.PlayTurnElim(playerTurn, game1.Players[randomVote])
			if err2 != nil {
				// _ = a
				fmt.Println(err2.Error(), " ", fmt.Sprint(game1.GameState))
			} else {
				fmt.Println("Player", playerTurn, "voted for", game1.Players[randomVote])
			}
		}

		finished, err := game1.IsGameFinished()
		if err != nil {
			fmt.Println("Error : ")
			fmt.Println(err)
			return
		}
		if finished.Winners != nil {
			fmt.Println(finished.Winners)
			break
		}

		time.Sleep(1000 * time.Millisecond)

		turn++

		fmt.Println("Next turn", turn)

	}

	fmt.Println(game1)
}
