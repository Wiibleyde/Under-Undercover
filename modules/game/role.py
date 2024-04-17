class Role:
    def __init__(self, name:str):
        self.name = name

ROLES = {
    "Normal": Role("Normal"),
    "Undercover": Role("Undercover"),
    "MrWhite": Role("Mr. White"),
    "UnSet": Role("UnSet")
}