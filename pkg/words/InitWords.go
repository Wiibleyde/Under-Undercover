package words

import (
	"encoding/csv"
	"math/rand"
	"os"
)

type CsvLine struct {
	NormalWord     string
	UndercoverWord string
}

type WordsType struct {
	NormalWords     []string
	UndercoverWords []string
}

var Words WordsType

func InitWords() {
	file, err := os.Open("words.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		Words.NormalWords = append(Words.NormalWords, record[0])
		Words.UndercoverWords = append(Words.UndercoverWords, record[1])
	}
}

func GetRandomWords() CsvLine {
	randomIndex := rand.Intn(len(Words.NormalWords))
	return CsvLine{
		NormalWord:     Words.NormalWords[randomIndex],
		UndercoverWord: Words.UndercoverWords[randomIndex],
	}
}
