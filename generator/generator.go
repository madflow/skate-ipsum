package generator

import (
	"encoding/json"
	"math/rand"
	"strings"

	"github.com/madflow/skate-ipsum/data"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomElement(elements []string) string {
	return elements[randInt(0, len(elements)-1)]
}

func Sentence(properNouns []string, phrases []string, numberOfWords int) string {
	wordList := make([]string, numberOfWords)

	useNounNext := true
	for i := 0; i < numberOfWords; i++ {
		if useNounNext {
			wordList[i] = randomElement(properNouns)
			useNounNext = false
		} else {
			wordList[i] = randomElement(phrases)
			useNounNext = true
		}
	}

	return strings.Join(wordList, " ") + "."
}

func Paragraph() string {
	numberOfSentences := randInt(3, 6)
	sentences := make([]string, numberOfSentences)
	for i := 0; i < numberOfSentences; i++ {
		sentences[i] = Sentence(data.ProperNouns, data.RandomPhrases, randInt(5, 10))
	}
	return strings.Join(sentences, " ")
}

func IpsumArray(numberOfParagraphs int) []string {
	paragraphs := make([]string, numberOfParagraphs)
	for i := 0; i < numberOfParagraphs; i++ {
		paragraphs[i] = Paragraph()
	}
	return paragraphs
}

func IpsumText(numberOfParagraphs int) string {
	return strings.Join(IpsumArray(numberOfParagraphs), "\n\n")
}

func IpsumJson(numberOfParagraphs int) string {
	parapgraphs := IpsumArray(numberOfParagraphs)
	builder := new(strings.Builder)
	err := json.NewEncoder(builder).Encode(parapgraphs)
	if err != nil {
		panic(err)
	}
	return builder.String()
}
