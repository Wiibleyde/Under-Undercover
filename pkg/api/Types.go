package api

import "game"

type Response struct {
	Message  Message          `json:"message"`
	GameData ResponseGameData `json:"gameData"`
	UserWord string           `json:"userWord"`
}

type Message struct {
	Information string `json:"information"`
	Warning     string `json:"warning"`
	Error       string `json:"error"`
}

type ResponseGameData struct {
	GameUuid   string                `json:"gameUuid"`
	Started    bool                  `json:"started"`
	Ended      bool                  `json:"ended"`
	PlayerTurn int                   `json:"playerTurn"`
	Players    []PlayerHidden        `json:"players"`
	PlaysDesc  []PlaysDescDataHidden `json:"playsDesc"`
	PlaysVote  []PlaysVoteDataHidden `json:"playsVote"`
	GameState  game.GameState        `json:"gameState"`
}

type PlayerHidden struct {
	Uuid       string `json:"uuid"`
	Pseudo     string `json:"pseudo"`
	Eliminated bool   `json:"eliminated"`
	Connected  bool   `json:"connected"`
}

type PlaysDescDataHidden struct {
	PlayerUuid string `json:"playerUuid"`
	Word       string `json:"word"`
}

type PlaysVoteDataHidden struct {
	PlayerUuid string `json:"playerUuid"`
	VoteUuid   string `json:"voteUuid"`
}
