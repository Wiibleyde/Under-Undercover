import csv
import random

class words:
    def __init__(self, normal:str, undercover:str):
        self.normal = normal
        self.undercover = undercover

wordsList: list[words] = []

def getWords():
    with open('words.csv', newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=',')
        for row in reader:
            wordsList.append(words(row[0], row[1]))

def getWord() -> words:
    return wordsList[random.randint(0, len(wordsList)-1)]