package anagram

import (
	"bufio"
	"bytes"
	"os"
)

func loadWords() ([]string, error) {
	f, err := os.ReadFile("testdata/words.txt")
	if err != nil {
		return []string{}, nil
	}
	s := bufio.NewScanner(bytes.NewBuffer(f))
	s.Split(bufio.ScanWords)
	var w []string
	for s.Scan() {
		w = append(w, s.Text())
	}
	return w, nil
}
