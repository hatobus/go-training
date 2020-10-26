package intset

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIntsetElems(t *testing.T) {
	s := &IntSet{
		words: []uint64{1, 2, 4},
	}

	out := s.Elems()

	if diff := cmp.Diff(out, []int{0, 65, 130}); diff != "" {
		t.Fatalf("invalid output diff = %v", diff)
	}
}
