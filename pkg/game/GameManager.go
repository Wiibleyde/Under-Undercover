package game

import (
	"errors"
	"fmt"
	"math/rand"
	"words"
)

func (g *Game) InitGame(host Player) {
	g.Host = host
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
	rand.Shuffle(len(g.Players), func(i, j int) {
		g.Players[i], g.Players[j] = g.Players[j], g.Players[i]
	})
}

func (g *Game) PlayTurnDesc(player Player, wordGiven string) error {
	if !g.GameState.DescriptionPhase {
		return errors.New(WrongAction.Message)
	}
	if player.Eliminated {
		g.SetNextPlayerTurn()
		return errors.New(PlayerEliminated.Message)
	}
	g.PlaysDesc = append(g.PlaysDesc, PlaysDescData{
		Turn:      g.PlayerTurn,
		Player:    player,
		WordGiven: wordGiven,
	})
	g.SetNextPlayerTurn()
	if len(g.PlaysDesc) == len(g.GetAlivePlayers()) {
		g.PlayerTurn = 0
		g.GameState.DescriptionPhase = false
		g.GameState.DiscussionPhase = true
		g.GameState.EliminationPhase = false
		g.PlaysDesc = []PlaysDescData{}
	}

	return nil
}

func (g *Game) PlayTurnDiscuss(player Player) error {
	if !g.GameState.DiscussionPhase {
		return errors.New(WrongAction.Message)
	}
	if g.Host.Uuid != player.Uuid {
		return errors.New(NotHost.Message)
	}
	g.PlayerTurn = 0
	g.GameState.DescriptionPhase = false
	g.GameState.DiscussionPhase = false
	g.GameState.EliminationPhase = true

	return nil
}

func (g *Game) PlayTurnElim(player Player, votedPlayer Player) error {
	if !g.GameState.EliminationPhase {
		return errors.New(WrongAction.Message)
	}
	if player.Eliminated {
		g.SetNextPlayerTurn()
		return errors.New(NotYourTurn.Message)
	}
	g.PlaysVote = append(g.PlaysVote, PlaysVoteData{
		Turn:   g.PlayerTurn,
		Player: player,
		Vote:   votedPlayer,
	})
	g.SetNextPlayerTurn()
	if len(g.PlaysVote) == len(g.GetAlivePlayers()) {
		var votes = make(map[string]int)
		for _, vote := range g.PlaysVote {
			votes[vote.Vote.Uuid]++
		}
		var maxVote int
		var eliminatedPlayer Player
		for uuid, vote := range votes {
			if vote > maxVote {
				maxVote = vote
				for _, player := range g.Players {
					if player.Uuid == uuid {
						eliminatedPlayer = player
					}
				}
			}
		}
		eliminatedPlayer.Eliminated = true
		for i, player := range g.Players {
			if player.Uuid == eliminatedPlayer.Uuid {
				g.Players[i] = eliminatedPlayer
			}
		}

		g.PlayerTurn = 0
		g.GameState.DescriptionPhase = true
		g.GameState.DiscussionPhase = false
		g.GameState.EliminationPhase = false
		g.PlaysVote = []PlaysVoteData{}
	}

	return nil
}

func (g *Game) IsGameFinished() (WinMessage, error) {
	var undercoverAlive = 0
	var normalAlive = 0
	var mrWhiteAlive = false
	for _, player := range g.Players {
		if player.Eliminated {
			continue
		}
		switch player.Role {
		case Undercover:
			undercoverAlive++
		case MrWhite:
			mrWhiteAlive = true
		case Normal:
			normalAlive++
		}
	}
	fmt.Println("Undercover alive: ", undercoverAlive, "Normal alive: ", normalAlive, "MrWhite alive: ", mrWhiteAlive)
	if undercoverAlive == 0 && !mrWhiteAlive && normalAlive == 0 {
		return WinMessage{WinRole: NotSet, Winners: []Player{}}, errors.New(NoWinnersError.Message)
	}
	if mrWhiteAlive && normalAlive <= 1 && undercoverAlive == 0 {
		mrWhitePlayer, err := g.GetPlayerByRole(MrWhite)
		if err == nil {
			return WinMessage{WinRole: MrWhite, Winners: []Player{mrWhitePlayer}}, nil
		}
		return WinMessage{}, errors.New(MrWhiteWinError.Message)
	}
	if undercoverAlive != 0 && normalAlive <= 1 && !mrWhiteAlive {
		undercoverPlayer, err := g.GetPlayerByRole(Undercover)
		if err == nil {
			return WinMessage{WinRole: Undercover, Winners: []Player{undercoverPlayer}}, nil
		}
		return WinMessage{}, errors.New(UndercoverWinError.Message)
	}
	if undercoverAlive == 0 && !mrWhiteAlive {
		normals := g.GetNormalPlayers()
		for i := range normals {
			if normals[i].Eliminated {
				normals = append(normals[:i], normals[i+1:]...)
			}
		}
		return WinMessage{WinRole: Normal, Winners: normals}, nil
	}
	return WinMessage{}, nil
}
