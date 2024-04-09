package game

import (
	"errors"

	"github.com/google/uuid"
)

var Games []Game

func (g *Game) AddPlayer(p Player) error {
	if g.Started {
		return errors.New(GameAlreadyStarted.Message)
	}
	g.Players = append(g.Players, p)
	return nil
}

func (g *Game) CreateGame() {
	g.Uuid = uuid.New().String()
	g.Started = false
	g.Ended = false
	g.GameState = GameState{
		DescriptionPhase: true,
		DiscussionPhase:  false,
		EliminationPhase: false,
	}
	g.PlayerTurn = 0
	g.Players = []Player{}
	g.Data = GameData{
		NormalWord:     "",
		UndercoverWord: "",
	}
}

func (g *Game) StartGame() error {
	if g.Started {
		return errors.New(GameAlreadyStarted.Message)
	}
	if len(g.Players) < 3 {
		return errors.New(NotEnoughPlayers.Message)
	}
	if g.Data.NormalWord == "" || g.Data.UndercoverWord == "" {
		return errors.New(WordsNotSet.Message)
	}
	for i, player := range g.Players {
		if player.Role == NotSet {
			return errors.New(RolesNotSet.Message)
		}
		// if player.Position == -1 {
		// 	return errors.New(PositionsNotSet.Message)
		// }
		g.Players[i] = player
	}

	g.Started = true
	g.GameState.DescriptionPhase = true
	g.GameState.DiscussionPhase = false
	g.GameState.EliminationPhase = false
	g.PlayerTurn = 0

	return nil
}

func (g *Game) GetPlayer(uuid string) (Player, error) {
	for _, p := range g.Players {
		if p.Uuid == uuid {
			return p, nil
		}
	}
	return Player{}, errors.New(PlayerNotFound.Message)
}

func (g *Game) GetPlayers() []Player {
	return g.Players
}

func (g *Game) GetGameState() GameState {
	return g.GameState
}

func (g *Game) GetGameData() GameData {
	return g.Data
}

func (g *Game) SetGameData(data GameData) {
	g.Data = data
}

func (g *Game) SetGameState(state GameState) {
	g.GameState = state
}

func (g *Game) SetPlayers(players []Player) {
	g.Players = players
}

func (g *Game) SetPlayer(p Player) {
	for i, player := range g.Players {
		if player.Uuid == p.Uuid {
			g.Players[i] = p
		}
	}
}

func (g *Game) RemovePlayer(uuid string) {
	for i, p := range g.Players {
		if p.Uuid == uuid {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
}

func (g *Game) GetPlayerByRole(role Role) (Player, error) {
	for _, p := range g.Players {
		if p.Role == role {
			return p, nil
		}
	}
	return Player{}, errors.New(PlayerNotFound.Message)
}

func (g *Game) GetNormalPlayers() []Player {
	var normalPlayers []Player
	for _, p := range g.Players {
		if p.Role == Normal {
			normalPlayers = append(normalPlayers, p)
		}
	}
	return normalPlayers
}

func (g *Game) SetPlayerByRole(role Role, p Player) {
	for i, player := range g.Players {
		if player.Role == role {
			g.Players[i] = p
		}
	}
}

func (g *Game) SetNormalPlayers(players []Player) {
	for i, p := range g.Players {
		if p.Role == Normal {
			g.Players[i] = players[i]
		}
	}
}

func (g *Game) GetNormalWord() string {
	return g.Data.NormalWord
}

func (g *Game) SetNormalWord(word string) {
	g.Data.NormalWord = word
}

func (g *Game) GetUndercoverWord() string {
	return g.Data.UndercoverWord
}

func (g *Game) SetUndercoverWord(word string) {
	g.Data.UndercoverWord = word
}

func (g *Game) NextGameState() {
	if g.GameState.DescriptionPhase {
		g.GameState.DescriptionPhase = false
		g.GameState.DiscussionPhase = true
	} else if g.GameState.DiscussionPhase {
		g.GameState.DiscussionPhase = false
		g.GameState.EliminationPhase = true
	} else if g.GameState.EliminationPhase {
		g.GameState.EliminationPhase = false
		g.GameState.DescriptionPhase = true
	}
}

func (g *Game) SetNextPlayerTurn() {
	var validPlayer bool
	for !validPlayer {
		if g.PlayerTurn == len(g.GetAlivePlayers()) {
			g.NextGameState()
			g.PlayerTurn = 0
		}
		if !g.Players[g.PlayerTurn].Eliminated {
			validPlayer = true
		}
		g.PlayerTurn++
	}
}

func (g *Game) GetNextPlayer() (Player, error) {
	for i := g.PlayerTurn + 1; i < len(g.Players); i++ {
		if !g.Players[i].Eliminated {
			g.PlayerTurn = i
			return g.Players[i], nil
		}
	}
	for i := 0; i < g.PlayerTurn; i++ {
		if !g.Players[i].Eliminated {
			g.PlayerTurn = i
			return g.Players[i], nil
		}
	}
	return Player{}, errors.New(NoNextPlayer.Message)
}

func (g *Game) GetAlivePlayers() []Player {
	var alivePlayers []Player
	for _, p := range g.Players {
		if !p.Eliminated {
			alivePlayers = append(alivePlayers, p)
		}
	}
	return alivePlayers
}

func GetGame(uuid string) (Game, error) {
	for _, g := range Games {
		if g.Uuid == uuid {
			return g, nil
		}
	}
	return Game{}, errors.New(GameNotFound.Message)
}
