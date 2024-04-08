package main

import (
	"fmt"
	"game"
	"math/rand"
	"words"
)

func main() {
	words.InitWords()

	player1 := game.Player{
		Uuid:     "1",
		Pseudo:   "player1",
		Position: -1,
		Role:     game.NotSet,
	}
	player2 := game.Player{
		Uuid:     "2",
		Pseudo:   "player2",
		Position: -1,
		Role:     game.NotSet,
	}
	player3 := game.Player{
		Uuid:     "3",
		Pseudo:   "player3",
		Position: -1,
		Role:     game.NotSet,
	}
	player4 := game.Player{
		Uuid:     "4",
		Pseudo:   "player4",
		Position: -1,
		Role:     game.NotSet,
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

	for {
		for a := range game1.Players {
			playerTurn, _ := game1.GetNextPlayer()
			err2 := game1.PlayTurnDesc(playerTurn, "word2")
			if err2 != nil {
				fmt.Println(game1.Players[a].Pseudo, " : ", err2.Error(), " ", fmt.Sprint(game1.GameState))
			}
		}

		err2 := game1.PlayTurnDiscuss(game1.Host)
		if err2 != nil {
			if err2.Error() == game.NotHost.Message {
				return
			}
			fmt.Println(err2, " ", fmt.Sprint(game1.GameState))
		}

		fmt.Println("Discussion phase over")
		fmt.Println("Elimination phase", game1.PlayerTurn)

		for a := range game1.Players {
			playerTurn, _ := game1.GetNextPlayer()
			randomVote := rand.Intn(len(game1.Players) - 1)
			err2 := game1.PlayTurnElim(playerTurn, game1.Players[randomVote])
			if err2 != nil {
				fmt.Println(game1.Players[a].Pseudo, " : ", err2.Error(), " ", fmt.Sprint(game1.GameState))
			}
		}

		finished, err := game1.IsGameFinished()
		if err != nil {
			fmt.Println("Error : ")
			fmt.Println(err)
			return
		}
		if finished.Winners != nil {
			fmt.Println(finished.WinRole)
			fmt.Println(finished.Winners)
			break
		}

		// time.Sleep(100 * time.Millisecond)

		fmt.Println("Next turn")

	}
}
