package intset

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddAll(t *testing.T) {
	intset := &IntSet{
		words: []uint64{1, 2, 3, 4, 5},
	}

	intset.AddAll()

	if diff := cmp.Diff(intset.words, []uint64{1, 2, 3, 4, 5}); diff != "" {
		t.Fatalf("invalid add all, diff: %v", diff)
	}

	intset.AddAll(6, 7, 8, 9, 0)

	if diff := cmp.Diff(intset.words, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}); diff != "" {
		t.Fatalf("invalid add all, diff: %v", diff)
	}
}
