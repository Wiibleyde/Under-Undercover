from modules.utils import csvUtil
from modules.game import gameObj, playerObj, roleObj

import flask
import datetime

app = flask.Flask(__name__)
games: list[gameObj.Game] = []

def getGame(gameId:str) -> gameObj.Game:
    for game in games:
        if game.uuid == gameId:
            return game
    return None

def getPlayer(playerId:str) -> playerObj.Player:
    for game in games:
        for player in game.players:
            if player.uuid == playerId:
                return player
    return None

def getPlayerInGame(playerId:str, gameId:str) -> playerObj.Player:
    for game in games:
        if game.uuid == gameId:
            for player in game.players:
                if player.uuid == playerId:
                    return player
    return None

def hideSensitiveDatas(game:gameObj.Game) -> gameObj.Game:
    newGame = game
    newGame.gameData = gameObj.GameData("nice", "try")
    for index in range(len(newGame.players)):
        newGame.players[index].role = roleObj.Role("etnoooon")
    return newGame

@app.route("/")
def index() -> str:
    return "Hello, World!"

@app.route("/createGame")
def createGame() -> str:
    newGame = gameObj.Game()
    games.append(newGame)
    return newGame.uuid

@app.route("/getCurrentGame")
def getCurrentGame() -> str:
    playerUuid = flask.request.cookies.get("playerUWUID")
    gameUuid = flask.request.cookies.get("gameUWUID")
    if playerUuid and gameUuid:
        game = getGame(gameUuid)
        player = getPlayerInGame(playerUuid, gameUuid)
        if game and player:
            hiddenGame = hideSensitiveDatas(game)
            return flask.jsonify(hiddenGame.__dict__())
    return "No game found"

@app.route("/joinGame/<gameId>/<pseudo>")
def joinGame(gameId:str, pseudo:str="Visitor"):
    playerUuid = flask.request.cookies.get("playerUWUID")
    if playerUuid:
        player = getPlayer(playerUuid)
        if player:
            return "Player already in a game"
    player = playerObj.Player(pseudo=pseudo)
    game = getGame(gameId)
    if game:
        game.addPlayer(player)
        if len(game.players) == 1:
            game.host = player
        resp = flask.make_response("Player joined")
        resp.set_cookie("playerUWUID", player.uuid, expires=datetime.datetime.now() + datetime.timedelta(hours=4))
        resp.set_cookie("gameUWUID", game.uuid, expires=datetime.datetime.now() + datetime.timedelta(hours=4))
        return resp
    return "Game not found"

if __name__ == "__main__":
    csvUtil.getWords()
    app.run(debug=True)