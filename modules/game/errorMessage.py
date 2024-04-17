class ErrorMessage:
    def __init__(self, code:int, message:str):
        self.code = code
        self.message = message

ERROR_MESSAGES = {
    "playerNotFound": ErrorMessage(1, "Joueur non trouvé"),
    "gameStarted": ErrorMessage(2, "Partie déjà commencée"),
    "alreadyInGame": ErrorMessage(3, "Joueur déjà dans la partie"),
    "notEnoughPlayers": ErrorMessage(4, "Pas assez de joueurs"),
    "notInitializedGame": ErrorMessage(5, "Partie non initialisée"),
    "incorrectGameStatus": ErrorMessage(6, "Statut de la partie incorrect"),
    "wrongActionStatus": ErrorMessage(7, "Action incorrecte (mauvais statut de la partie)"),
    "notYourTurn": ErrorMessage(8, "Action incorrecte (pas votre tour)"),
    "hostOnly": ErrorMessage(9, "Action incorrecte (seul l'hôte peut effectuer cette action)"),
}

