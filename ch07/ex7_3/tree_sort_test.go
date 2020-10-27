package ex7_3

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTreeSort(t *testing.T) {
	root := &Tree{value: 0}

	root = add(root, 1)

	rs := root.String()

	if diff := cmp.Diff(rs, "[0 1]"); diff != "" {
		t.Fatalf("invalid string diff = %v", diff)
	}

	root = add(root, 5)

	rs = root.String()

	if diff := cmp.Diff(rs, "[0 1 5]"); diff != "" {
		t.Fatalf("invalid string diff = %v", diff)
	}

	root = add(root, 2)

	rs = root.String()

	if diff := cmp.Diff(rs, "[0 1 2 5]"); diff != "" {
		t.Fatalf("invalid string diff = %v", diff)
	}
}
