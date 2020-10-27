package counter

import (
	"bufio"
	"strings"
)

type LineCounter struct {
	lines int
}

func NewLineCounter() *LineCounter {
	return &LineCounter{}
}

func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lc.lines++
	}
	return len(p), nil
}

func (lc *LineCounter) Lines() int {
	return lc.lines
}
