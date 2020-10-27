package counter

import (
	"bufio"
	"strings"
)

type WordCounter struct {
	words int
}

func NewWordCounter() *WordCounter {
	return &WordCounter{}
}

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wc.words++
	}

	return len(p), nil
}

func (wc *WordCounter) Words() int {
	return wc.words
}
