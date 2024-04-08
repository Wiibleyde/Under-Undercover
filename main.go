package main

import (
	"game"
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

	game1.InitGame()

	for _, player := range game1.Players {
		println(player.Role, player.Pseudo, player.Position)
	}
}
