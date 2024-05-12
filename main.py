from modules.utils import csvUtil
from modules.game import gameObj, playerObj, roleObj
from modules.game.roleObj import ROLES

import flask
from flask_cors import CORS
import secrets
import datetime
import uuid

app = flask.Flask(__name__)
app.secret_key = secrets.token_urlsafe(16)
CORS(app, supports_credentials=True, resources={r"/*": {"origins": "*"}})
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

@app.route("/", methods=["GET"])
@app.route("/status", methods=["GET"])
def status() -> str:
    return flask.jsonify({"status": "ok"})

@app.route("/createGame", methods=["POST"])
def createGame() -> str:
    newGame = gameObj.Game(uuid=str(uuid.uuid4()))
    games.append(newGame)
    return newGame.uuid

@app.route("/getCurrentGame", methods=["GET"])
def getCurrentGame() -> str:
    playerUuid = flask.request.cookies.get("playerUWUID")
    gameUuid = flask.request.cookies.get("gameUWUID")
    if playerUuid and gameUuid:
        game = getGame(gameUuid)
        player = getPlayerInGame(playerUuid, gameUuid)
        if game and player:
            return flask.jsonify(game.__dict__())
    return flask.jsonify({"error": "Player or game not found"})

@app.route("/joinGame", methods=["POST"])
def joinGame():
    gameId = flask.request.args.get("gameId")
    pseudo = flask.request.args.get("pseudo")
    playerUuid = flask.request.cookies.get("playerUWUID")
    if playerUuid:
        player = getPlayer(playerUuid)
        if player:
            return flask.jsonify({"error": "Player already in a game"})
    player = playerObj.Player(uuid=str(uuid.uuid4()), pseudo=pseudo)
    game = getGame(gameId)
    if game:
        errMessage = game.addPlayer(player)
        if errMessage:
            return flask.jsonify({"error": errMessage.message})
        if len(game.players) == 1:
            game.host = player
        resp = flask.make_response(flask.jsonify(game.__dict__()))
        resp.set_cookie("playerUWUID", player.uuid, expires=datetime.datetime.now() + datetime.timedelta(hours=1))
        resp.set_cookie("gameUWUID", game.uuid, expires=datetime.datetime.now() + datetime.timedelta(hours=1))
        return resp
    return flask.jsonify({"error": "Game not found"})

@app.route("/leaveGame", methods=["POST"])
def leaveGame():
    playerUuid = flask.request.cookies.get("playerUWUID")
    gameUuid = flask.request.cookies.get("gameUWUID")
    game = getGame(gameUuid)
    if game:
        player = getPlayerInGame(playerUuid, gameUuid)
        if player:
            errMessage = game.removePlayer(player)
            if errMessage:
                return errMessage.message
            resp = flask.make_response(flask.jsonify({"status": "Player left"}))
            resp.set_cookie("playerUWUID", "", expires=0)
            return resp
        return flask.jsonify({"error": "Player not found"})
    return flask.jsonify({"error": "Game not found"})

@app.route("/startGame", methods=["POST"])
def startGame():
    playerUuid = flask.request.cookies.get("playerUWUID")
    gameUuid = flask.request.cookies.get("gameUWUID")
    game = getGame(gameUuid)
    if game:
        player = getPlayerInGame(playerUuid, gameUuid)
        errMessage = game.startGame(player)
        if errMessage is None:
            return flask.jsonify({"status": "Game started"})
        return flask.jsonify({"error": errMessage.message})
    return flask.jsonify({"error": "Game not found"})

@app.route("/playDescTurn", methods=["POST"])
def playDescTurn():
    playerUuid = flask.request.cookies.get("playerUWUID")
    gameUuid = flask.request.cookies.get("gameUWUID")
    game = getGame(gameUuid)
    if game:
        player = getPlayerInGame(playerUuid, gameUuid)
        if player:
            desc = flask.request.args.get("desc")
            errMessage = game.playDesc(player, desc)
            if errMessage is None:
                return flask.jsonify({"status": "Description played"})
            return flask.jsonify({"error": errMessage.message})
        return flask.jsonify({"error": "Player not found"})
    return flask.jsonify({"error": "Game not found"})

@app.route("/playVoteTurn", methods=["POST"])
def playVoteTurn():
    playerUuid = flask.request.cookies.get("playerUWUID")
    gameUuid = flask.request.cookies.get("gameUWUID")
    game = getGame(gameUuid)
    if game:
        player = getPlayerInGame(playerUuid, gameUuid)
        if player:
            vote = flask.request.args.get("vote")
            errMessage = game.playVote(player, vote)
            if errMessage is None:
                return flask.jsonify({"status": "Vote played"})
            return flask.jsonify({"error": errMessage.message})
        return flask.jsonify({"error": "Player not found"})
    return flask.jsonify({"error": "Game not found"})

@app.route("/playDiscussionTurn", methods=["POST"])
def playDiscussionTurn():
    playerUuid = flask.request.cookies.get("playerUWUID")
    gameUuid = flask.request.cookies.get("gameUWUID")
    game = getGame(gameUuid)
    if game:
        player = getPlayerInGame(playerUuid, gameUuid)
        if player:
            errMessage = game.playDiscussion(player)
            if errMessage is None:
                return flask.jsonify({"status": "Discussion played"})
            return flask.jsonify({"error": errMessage.message})
        return flask.jsonify({"error": "Player not found"})
    return flask.jsonify({"error": "Game not found"})

if __name__ == "__main__":
    csvUtil.getWords()
    app.run(debug=True)