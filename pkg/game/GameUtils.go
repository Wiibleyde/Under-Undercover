package game

import (
	"errors"
	"logger"
	"time"

	"github.com/google/uuid"
)

var Games []Game

func (g *Game) AddPlayer(p Player) error {
	if g.Started {
		return errors.New(GameAlreadyStarted.Message)
	}
	for _, player := range g.Players {
		if player.Uuid == p.Uuid {
			return errors.New(PlayerAlreadyInGame.Message)
		}
	}
	g.Players = append(g.Players, p)
	if len(g.Players) == 1 {
		g.Host = p
	}
	UpdateGame(*g)
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
	Games = append(Games, *g)
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

	UpdateGame(*g)

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
	UpdateGame(*g)
}

func (g *Game) SetGameState(state GameState) {
	g.GameState = state
	UpdateGame(*g)
}

func (g *Game) SetPlayers(players []Player) {
	g.Players = players
	UpdateGame(*g)
}

func (g *Game) SetPlayer(p Player) {
	for i, player := range g.Players {
		if player.Uuid == p.Uuid {
			g.Players[i] = p
		}
	}
	UpdateGame(*g)
}

func (g *Game) RemovePlayer(uuid string) error {
	var edited bool
	for i, p := range g.Players {
		if p.Uuid == uuid {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			edited = true
		}
	}
	if !edited {
		return errors.New(PlayerNotFound.Message)
	}
	UpdateGame(*g)
	return nil
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
	UpdateGame(*g)
}

func (g *Game) SetNormalPlayers(players []Player) {
	for i, p := range g.Players {
		if p.Role == Normal {
			g.Players[i] = players[i]
		}
	}
	UpdateGame(*g)
}

func (g *Game) GetNormalWord() string {
	return g.Data.NormalWord
}

func (g *Game) SetNormalWord(word string) {
	g.Data.NormalWord = word
	UpdateGame(*g)
}

func (g *Game) GetUndercoverWord() string {
	return g.Data.UndercoverWord
}

func (g *Game) SetUndercoverWord(word string) {
	g.Data.UndercoverWord = word
	UpdateGame(*g)
}

func (g *Game) SetNextGameState() {
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
	UpdateGame(*g)
}

func (g *Game) SetNextPlayerTurn() {
	var validPlayer bool
	for !validPlayer {
		if g.PlayerTurn == len(g.GetAlivePlayers()) {
			g.SetNextGameState()
			g.PlayerTurn = 0
		}
		if !g.Players[g.PlayerTurn].Eliminated {
			validPlayer = true
		}
		g.PlayerTurn++
	}
	UpdateGame(*g)
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

func (g *Game) GetPlayerWord(player Player) string {
	for _, p := range g.Players {
		if p.Uuid == player.Uuid {
			role := p.Role
			if role == Normal {
				return g.Data.NormalWord
			}
			if role == Undercover {
				return g.Data.UndercoverWord
			}
		}
	}
	return ""
}

func GetGame(uuid string) (Game, error) {
	for _, g := range Games {
		if g.Uuid == uuid {
			return g, nil
		}
	}
	return Game{}, errors.New(GameNotFound.Message)
}

func UpdateGame(game Game) {
	for i, g := range Games {
		if g.Uuid == game.Uuid {
			game.LastUpdate = time.Now()
			Games[i] = game
			return
		}
	}
	logger.ErrorLogger.Println(GameNotFound.Message)
}
