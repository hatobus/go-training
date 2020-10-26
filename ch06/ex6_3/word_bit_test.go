package intset

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWordIntersectWith(t *testing.T) {
	t.Parallel()

	s := &IntSet{
		words: []uint64{0, 2, 4},
	}

	s2 := &IntSet{
		words: []uint64{1, 2, 3},
	}

	s.IntersectWith(s2)

	if diff := cmp.Diff(s.words, []uint64{0, 2, 0}); diff != "" {
		t.Fatalf("ivalid output diff = %v", diff)
	}
}

func TestDifferenceWith(t *testing.T) {
	t.Parallel()

	s := &IntSet{
		words: []uint64{0, 2, 4},
	}

	s2 := &IntSet{
		words: []uint64{1, 2, 3},
	}

	s.DifferenceWith(s2)

	if diff := cmp.Diff(s.words, []uint64{0, 0, 4}); diff != "" {
		t.Fatalf("invalid output diff = %v", diff)
	}
}

func TestSymmetricDifference(t *testing.T) {
	t.Parallel()

	s := &IntSet{
		words: []uint64{0, 2, 4},
	}

	s2 := &IntSet{
		words: []uint64{1, 2, 3},
	}

	s.SymmetricDifference(s2)

	if diff := cmp.Diff(s.words, []uint64{1, 0, 7}); diff != "" {
		t.Fatalf("invald output diff = %v", diff)
	}
}
