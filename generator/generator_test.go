package generator

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestRandInt(t *testing.T) {
	min := 1
	max := 10
	result := randInt(min, max)
	if result < min || result > max {
		t.Errorf("randInt result %d not within range [%d, %d]", result, min, max)
	}
}

func TestRandomElement(t *testing.T) {
	elements := []string{"a", "b", "c", "d"}
	result := randomElement(elements)
	if !contains(elements, result) {
		t.Errorf("randomElement result %s not found in elements", result)
	}
}

func TestSentence(t *testing.T) {
	properNouns := []string{"Alice", "Bob", "Charlie"}
	phrases := []string{"lorem", "ipsum", "dolor", "sit"}
	numberOfWords := 10
	result := Sentence(properNouns, phrases, numberOfWords)
	words := strings.Fields(result)
	if len(words) != numberOfWords {
		t.Errorf("Sentence has %d words, expected %d", len(words), numberOfWords)
	}
}

func TestParagraph(t *testing.T) {
	result := Paragraph()
	if !strings.Contains(result, ".") {
		t.Errorf("Paragraph does not end with a period: %s", result)
	}
}

func TestIpsumArray(t *testing.T) {
	numberOfParagraphs := 3
	result := IpsumArray(numberOfParagraphs)
	if len(result) != numberOfParagraphs {
		t.Errorf("IpsumArray has %d paragraphs, expected %d", len(result), numberOfParagraphs)
	}
}

func TestIpsumText(t *testing.T) {
	numberOfParagraphs := 3
	result := IpsumText(numberOfParagraphs)
	paragraphs := strings.Split(result, "\n")
	if len(paragraphs) != numberOfParagraphs {
		t.Errorf("IpsumText has %d paragraphs, expected %d", len(paragraphs), numberOfParagraphs)
	}
}

func TestIpsumJson(t *testing.T) {
	numberOfParagraphs := 3
	result := IpsumJson(numberOfParagraphs)
	var paragraphs []string
	err := json.Unmarshal([]byte(result), &paragraphs)
	if err != nil {
		t.Errorf("Error parsing JSON: %s", err)
	}
	if len(paragraphs) != numberOfParagraphs {
		t.Errorf("IpsumJson has %d paragraphs, expected %d", len(paragraphs), numberOfParagraphs)
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
