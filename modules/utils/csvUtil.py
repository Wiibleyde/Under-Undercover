import csv
import random

class Words:
    def __init__(self, normal:str, undercover:str):
        self.normal = normal
        self.undercover = undercover

wordsList: list[Words] = []

def getWords():
    with open('words.csv', newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=',')
        for row in reader:
            wordsList.append(Words(row[0], row[1]))

def getWord() -> Words:
    return wordsList[random.randint(0, len(wordsList)-1)]