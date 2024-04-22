class WinMessage:
    def __init__(self, code:int, message:str):
        self.code = code
        self.message = message

    def __dict__(self) -> dict:
        return {
            "code": self.code,
            "message": self.message
        }

WIN_MESSAGES = {
    "NormalWin": WinMessage(1, "Les civils ont gagné, tous les undercover et Mr. White ont été éliminés"),
    "UndercoverWin": WinMessage(2, "Les undercover ont gagné, les civils n'ont pas réussi à les éliminer"),
    "MrWhiteWin": WinMessage(3, "Mr. White a gagné, il a réussi à rester caché")
}