package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBreadsFirst(t *testing.T) {
	f := func(s string) []string {
		return prereqs[s]
	}

	out := breadthFirst(f, prereqs["algorithms"])

	expect := []string{"discrete math", "intro to programming"}

	if diff := cmp.Diff(out, expect); diff != "" {
		t.Fatalf("invalid output diff: %v", diff)
	}
}
