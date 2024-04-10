package api

import "game"

type Response struct {
	Message  Message   `json:"message"`
	GameData game.Game `json:"gameData"`
}

type Message struct {
	Information string `json:"information"`
	Warning     string `json:"warning"`
	Error       string `json:"error"`
}
