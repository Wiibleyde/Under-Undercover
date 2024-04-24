import uuid

from .roleObj import Role


class Player:
    def __init__(self, uuid:str, pseudo:str="Visitor", role:Role=None, eliminated:bool=False, connected:bool=True):
        self.uuid = uuid
        self.pseudo = pseudo
        self.role = role
        self.eliminated = eliminated
        self.connected = connected

    def __dict__(self) -> dict:
        return {
            "uuid": self.uuid,
            "pseudo": self.pseudo,
            "role": self.role.__dict__() if self.role else None,
            "eliminated": self.eliminated,
            "connected": self.connected
        }

    def setRole(self, role:Role):
        self.role = role

    def setEliminated(self, eliminated:bool):
        self.eliminated = eliminated

    def setConnected(self, connected:bool):
        self.connected = connected