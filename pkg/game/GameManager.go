package game

import (
	"math/rand"
	"words"
)

func (g *Game) InitGame() {
	g.SelectWords()
	g.AttributeRoles()
	g.ChoosePlayerOrder()
}

func (g *Game) SelectWords() {
	selectedWords := words.GetRandomWords()
	g.Data.NormalWord = selectedWords.NormalWord
	g.Data.UndercoverWord = selectedWords.UndercoverWord
}

func (g *Game) AttributeRoles() {
	roles := make([]Role, len(g.Players))
	if len(g.Players) > 3 {
		roles[0] = Undercover
		roles[1] = MrWhite
		for i := 2; i < len(roles); i++ {
			roles[i] = Normal
		}
	} else {
		roles[0] = Undercover
		for i := 1; i < len(roles); i++ {
			roles[i] = Normal
		}
	}

	rand.Shuffle(len(roles), func(i, j int) {
		roles[i], roles[j] = roles[j], roles[i]
	})

	for i, player := range g.Players {
		player.Role = roles[i]
		g.Players[i] = player
	}
}

func (g *Game) ChoosePlayerOrder() {
	for i := range g.Players {
		g.Players[i].Position = i
	}
	rand.Shuffle(len(g.Players), func(i, j int) {
		g.Players[i], g.Players[j] = g.Players[j], g.Players[i]
	})
}
