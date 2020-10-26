package intset

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntSet(t *testing.T) {
	t.Parallel()
	intset := &IntSet{
		words: []uint{1, 2, 3, 4, 5},
	}

	if diff := cmp.Diff(intset.Len(), 5); diff != "" {
		t.Fatalf("invalid inset length, diff: %v", diff)
	}

	intset.Remove(5)

	if diff := cmp.Diff(intset.words, []uint{1, 2, 3, 4}); diff != "" {
		t.Fatalf("invalid words, diff = %v", diff)
	}

	newIntSet := intset.Copy()

	if diff := cmp.Diff(intset.words, newIntSet.words); diff != "" {
		t.Fatalf("intset Copy invalid output, diff = %v", diff)
	}
}

func TestAddAll(t *testing.T) {
	t.Parallel()
	intset := &IntSet{
		words: []uint{1, 2, 3, 4, 5},
	}

	intset.AddAll()

	if diff := cmp.Diff(intset.words, []uint{1, 2, 3, 4, 5}); diff != "" {
		t.Fatalf("invalid add all, diff: %v", diff)
	}

	intset.AddAll(6, 7, 8, 9, 0)

	if diff := cmp.Diff(intset.words, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}); diff != "" {
		t.Fatalf("invalid add all, diff: %v", diff)
	}
}

func TestWordIntersectWith(t *testing.T) {
	t.Parallel()

	s := &IntSet{
		words: []uint{0, 2, 4},
	}

	s2 := &IntSet{
		words: []uint{1, 2, 3},
	}

	s.IntersectWith(s2)

	if diff := cmp.Diff(s.words, []uint{0, 2, 0}); diff != "" {
		t.Fatalf("ivalid output diff = %v", diff)
	}
}

func TestDifferenceWith(t *testing.T) {
	t.Parallel()

	s := &IntSet{
		words: []uint{0, 2, 4},
	}

	s2 := &IntSet{
		words: []uint{1, 2, 3},
	}

	s.DifferenceWith(s2)

	if diff := cmp.Diff(s.words, []uint{0, 0, 4}); diff != "" {
		t.Fatalf("invalid output diff = %v", diff)
	}
}

func TestSymmetricDifference(t *testing.T) {
	t.Parallel()

	s := &IntSet{
		words: []uint{0, 2, 4},
	}

	s2 := &IntSet{
		words: []uint{1, 2, 3},
	}

	s.SymmetricDifference(s2)

	if diff := cmp.Diff(s.words, []uint{1, 0, 7}); diff != "" {
		t.Fatalf("invald output diff = %v", diff)
	}
}

func TestIntsetElems(t *testing.T) {
	t.Parallel()
	s := &IntSet{
		words: []uint{1, 2, 4},
	}

	out := s.Elems()

	if diff := cmp.Diff(out, []int{0, 65, 130}); diff != "" {
		t.Fatalf("invalid output diff = %v", diff)
	}
}
