class Role:
    def __init__(self, name:str):
        self.name = name

    def __dict__(self) -> dict:
        return {
            "name": self.name
        }

ROLES = {
    "Normal": Role("Normal"),
    "Undercover": Role("Undercover"),
    "MrWhite": Role("Mr. White"),
    "UnSet": Role("UnSet")
}