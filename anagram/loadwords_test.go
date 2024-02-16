package anagram

import "testing"

func TestAnagramLoadWords(t *testing.T) {
	words, err := loadWords()
	if err != nil {
		t.Fatalf("error loading words: %v", err)
	}

	if words[0] != "acrobat" {
		t.Errorf(`expected "acrobat", got %q`, words[0])
	}
	if words[len(words)-1] != "yes" {
		t.Errorf(`expected "yes", got %q`, words[len(words)-1])
	}
}
