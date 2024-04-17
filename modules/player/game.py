from player import Player
from role import Role, ROLES

import uuid
import time

class GameData:
    def __init__(self, normalWord:str, undercoverWord:str):
        self.normalWord = normalWord
        self.undercoverWord = undercoverWord

class GameState:
    def __init__(self, state:int):
        self.state = state

    def getState(self):
        if self.state == 0:
            return "description"
        elif self.state == 1:
            return "discussion"
        elif self.state == 2:
            return "vote"
        else:
            return "unknown"

class DescPlayData:
    def __init__(self, player:Player, word:str):
        self.player = player
        self.word = word

class VoteData:
    def __init__(self, player:Player, targetPlayer:Player):
        self.player = player
        self.targetPlayer = targetPlayer

class Game:
    def __init__(self, uuid:str=str(uuid.uuid4()), started:bool=False, ended:bool=False, host:Player=Player(), players:list[Player]=[], gameData:GameData=GameData("",""), gameState:GameState=GameState(-1), playerTurn:int=0, descPlayData:list[DescPlayData]=None, voteData:list[VoteData]=None, lastUpdate:int=time.time()):
        self.uuid = uuid
        self.started = started
        self.ended = ended
        self.host = host
        self.players = players
        self.gameData = gameData
        self.gameState = gameState
        self.playerTurn = playerTurn
        self.descPlayData = descPlayData
        self.voteData = voteData
        self.lastUpdate = lastUpdate

    def getPlayer(self, uuid:str) -> Player:
        for player in self.players:
            if player.uuid == uuid:
                return player
        return None

    def addPlayer(self, player:Player) -> bool:
        if not self.started:
            return False
        if player in self.players:
            return False
        self.players.append(player)
        return True

    def removePlayer(self, player:Player) -> bool:
        if player not in self.players:
            return False
        self.players.remove(player)
        return True

    def startGame(self) -> bool:
        if self.started:
            return False
        if len(self.players) < 3:
            return False
        if self.gameData.normalWord == "" or self.gameData.undercoverWord == "":
            return False
        for player in self.players:
            if player.role == None or player.role == ROLES["UnSet"]:
                return False
        self.started = True
        self.gameState = GameState(0)
        return True

    def endGame(self) -> bool:
        if not self.started:
            return False
        self.ended = True
        return True

    def update(self) -> bool:
        self.lastUpdate = time.time()
        return True

    def nextGameState(self) -> str:
        if self.gameState.state == 0:
            self.gameState = GameState(1)
        elif self.gameState.state == 1:
            self.gameState = GameState(2)
        elif self.gameState.state == 2:
            self.gameState = GameState(0)
        return self.gameState.getState()

    def getPlayerTurn(self) -> Player:
        player = self.players[self.playerTurn]
        if self.gameState.getState() == "vote":
            if self.voteData != None:
                for vote in self.voteData:
                    if vote.player == player:
                        self.playerTurn += 1
                        player = self.players[self.playerTurn]
                    else:
                        return player
        elif self.gameState.getState() == "description":
            if self.descPlayData != None:
                for desc in self.descPlayData:
                    if desc.player == player:
                        self.playerTurn += 1
                        player = self.players[self.playerTurn]
                    else:
                        return player
        elif self.gameState.getState() == "discussion":
            return self.host
        else:
            return None

    def playDesc(self, player:Player, word:str) -> bool:
        if self.gameState.getState() != "description":
            return False
        if player != self.getPlayerTurn():
            return False
        if self.descPlayData == None:
            self.descPlayData = []
        self.descPlayData.append(DescPlayData(player, word))
        if len(self.descPlayData) == len(self.players):
            self.nextGameState()
        return True

    def playVote(self, player:Player, targetPlayer:Player) -> bool:
        if self.gameState.getState() != "vote":
            return False
        if player != self.getPlayerTurn():
            return False
        if self.voteData == None:
            self.voteData = []
        self.voteData.append(VoteData(player, targetPlayer))
        if len(self.voteData) == len(self.players):
            votes = {}
            for vote in self.voteData:
                if vote.targetPlayer in votes:
                    votes[vote.targetPlayer] += 1
                else:
                    votes[vote.targetPlayer] = 1
            maxVotes = 0
            maxPlayer = None
            for player, votes in votes.items():
                if votes > maxVotes:
                    maxVotes = votes
                    maxPlayer = player
            maxPlayer.setEliminated(True)
            self.nextGameState()
        return True

    def playDiscussion(self, player:Player) -> bool:
        if self.gameState.getState() != "discussion":
            return False
        if player != self.host:
            return False
        self.nextGameState()
        return True