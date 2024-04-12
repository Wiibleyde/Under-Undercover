package game

import "time"

type Role string

const (
	Undercover Role = "undercover"
	MrWhite    Role = "mrWhite"
	Normal     Role = "normal"
	NotSet     Role = "notSet"
)

type Game struct {
	Uuid       string          `json:"uuid"`
	Started    bool            `json:"started"`
	Ended      bool            `json:"ended"`
	GameState  GameState       `json:"gameState"`
	PlayerTurn int             `json:"playerTurn"`
	Host       Player          `json:"host"`
	Players    []Player        `json:"players"`
	Data       GameData        `json:"data"`
	PlaysDesc  []PlaysDescData `json:"plays"`
	PlaysVote  []PlaysVoteData `json:"votes"`
	LastUpdate time.Time       `json:"lastUpdate"`
}

type Player struct {
	Uuid       string `json:"uuid"`
	Pseudo     string `json:"pseudo"`
	Role       Role   `json:"role"`
	Eliminated bool   `json:"eliminated"`
	Connected  bool   `json:"connected"`
}

type GameData struct {
	NormalWord     string `json:"normalWord"`
	UndercoverWord string `json:"undercoverWord"`
}

type PlaysDescData struct {
	Turn      int    `json:"turn"`
	Player    Player `json:"player"`
	WordGiven string `json:"wordGiven"`
}

type PlaysVoteData struct {
	Turn   int    `json:"turn"`
	Player Player `json:"player"`
	Vote   Player `json:"vote"`
}

type GameState struct {
	DescriptionPhase bool `json:"descriptionPhase"`
	DiscussionPhase  bool `json:"discussionPhase"`
	EliminationPhase bool `json:"eliminationPhase"`
}

type GameError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	NotEnoughPlayers    = GameError{Code: 1, Message: "Pas assez de joueurs"}
	WordsNotSet         = GameError{Code: 2, Message: "Les mots n'ont pas été définis"}
	RolesNotSet         = GameError{Code: 3, Message: "Un joueur ou plusieurs n'ont pas de rôle"}
	PositionsNotSet     = GameError{Code: 4, Message: "Un joueur ou plusieurs n'ont pas d'ordre de passage"}
	GameAlreadyStarted  = GameError{Code: 5, Message: "La partie a déjà commencé"}
	GameNotFound        = GameError{Code: 6, Message: "La partie n'a pas été trouvée"}
	PlayerNotFound      = GameError{Code: 7, Message: "Le joueur n'a pas été trouvé"}
	PlayerEliminated    = GameError{Code: 8, Message: "Le joueur a été éliminé"}
	NotYourTurn         = GameError{Code: 8, Message: "Ce n'est pas votre tour"}
	WrongAction         = GameError{Code: 9, Message: "Action incorrecte"}
	NotHost             = GameError{Code: 10, Message: "Vous n'êtes pas l'hôte de la partie"}
	MrWhiteWinError     = GameError{Code: 11, Message: "Mr White a gagné mais le joueur est introuvable"}
	UndercoverWinError  = GameError{Code: 12, Message: "Les Undercovers ont gagné mais le joueur est introuvable"}
	NoWinnersError      = GameError{Code: 13, Message: "Aucun gagnant trouvé"}
	NoNextPlayer        = GameError{Code: 14, Message: "Aucun joueur suivant trouvé"}
	PlayerAlreadyInGame = GameError{Code: 15, Message: "Le joueur est déjà dans la partie"}
	AlreadyPlayed       = GameError{Code: 16, Message: "Le joueur a déjà joué"}
)

type WinMessage struct {
	WinRole Role `json:"winRole"`
	Winners []Player
}
