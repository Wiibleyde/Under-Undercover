import uuid
import json

from role import Role


class Player:
    def __init__(self, uuid:str=str(uuid.uuid4()), pseudo:str="Visitor", role:Role=None, eliminated:bool=False, connected:bool=True):
        self.uuid = uuid
        self.pseudo = pseudo
        self.role = role
        self.eliminated = eliminated
        self.connected = connected

    def setRole(self, role:Role):
        self.role = role

    def setEliminated(self, eliminated:bool):
        self.eliminated = eliminated

    def setConnected(self, connected:bool):
        self.connected = connected