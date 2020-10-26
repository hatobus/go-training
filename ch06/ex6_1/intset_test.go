package intset

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntSet(t *testing.T) {
	intset := &IntSet{
		words: []uint64{1, 2, 3, 4, 5},
	}

	if diff := cmp.Diff(intset.Len(), 5); diff != "" {
		t.Fatalf("invalid inset length, diff: %v", diff)
	}

	intset.Remove(5)

	if diff := cmp.Diff(intset.words, []uint64{1, 2, 3, 4}); diff != "" {
		t.Fatalf("invalid words, diff = %v", diff)
	}

	newIntSet := intset.Copy()

	if diff := cmp.Diff(intset.words, newIntSet.words); diff != "" {
		t.Fatalf("intset Copy invalid output, diff = %v", diff)
	}
}
