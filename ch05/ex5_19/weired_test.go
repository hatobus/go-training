package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWeired(t *testing.T) {
	out := Weird()

	if diff := cmp.Diff(out, "hello"); diff != "" {
		t.Fatalf("invalid output diff: %v", diff)
	}
}
