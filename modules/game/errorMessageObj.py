class ErrorMessage:
    def __init__(self, code:int, message:str):
        self.code = code
        self.message = message

    def __dict__(self) -> dict:
        return {
            "code": self.code,
            "message": self.message
        }

ERROR_MESSAGES = {
    "GameAlreadyStarted": ErrorMessage(1, "Partie déjà commencée"),
    "PlayerAlreadyInGame": ErrorMessage(2, "Joueur déjà dans la partie"),
    "PlayerNotInGame": ErrorMessage(3, "Joueur inexistant dans la partie"),
    "NotEnoughPlayers": ErrorMessage(4, "Pas assez de joueurs"),
    "GameNotStarted": ErrorMessage(5, "Partie non commencée"),
    "NotRightState": ErrorMessage(6, "Pas le bon état de jeu"),
    "NotYourTurn": ErrorMessage(7, "Pas votre tour"),
    "HostOnly": ErrorMessage(8, "Seul l'hôte peut faire cette action"),
}

