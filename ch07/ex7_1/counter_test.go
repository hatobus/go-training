package counter

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLineCounter(t *testing.T) {
	t.Parallel()

	lc := NewLineCounter()

	testString := fmt.Sprintf("a\nb\nc\n")

	n, err := lc.Write([]byte(testString))
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(len([]byte(testString)), n); diff != "" {
		t.Fatalf("invalid return length, diff = %v", diff)
	}

	if diff := cmp.Diff(lc.Lines(), 3); diff != "" {
		t.Fatalf("invalid line length diff = %v", diff)
	}
}

func TestLinesCounter(t *testing.T) {
	t.Parallel()

	wc := NewWordCounter()

	testWords := "this is a test string"

	n, err := wc.Write([]byte(testWords))
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(len([]byte(testWords)), n); diff != "" {
		t.Fatalf("invalid words length diff = %v", diff)
	}

	if diff := cmp.Diff(wc.Words(), 5); diff != "" {
		t.Fatalf("invalid word length diff = %v", diff)
	}
}
