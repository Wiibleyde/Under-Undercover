package game

type Role string

const (
	Undercover Role = "undercover"
	MrWhite    Role = "mrWhite"
	Normal     Role = "normal"
	NotSet     Role = "notSet"
)

type Player struct {
	Uuid     string `json:"uuid"`
	Pseudo   string `json:"pseudo"`
	Position int    `json:"position"`
	Role     Role   `json:"role"`
}

type GameData struct {
	NormalWord     string `json:"normalWord"`
	UndercoverWord string `json:"undercoverWord"`
}

type GameState struct {
	DescriptionPhase bool `json:"descriptionPhase"`
	DiscussionPhase  bool `json:"discussionPhase"`
	EliminationPhase bool `json:"eliminationPhase"`
}

type Game struct {
	Uuid       string    `json:"uuid"`
	Started    bool      `json:"started"`
	Ended      bool      `json:"ended"`
	GameState  GameState `json:"gameState"`
	PlayerTurn int       `json:"playerTurn"`
	Players    []Player  `json:"players"`
	Data       GameData  `json:"data"`
}
