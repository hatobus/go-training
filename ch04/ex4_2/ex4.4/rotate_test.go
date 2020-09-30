package rotate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRotate(t *testing.T) {
	inp := []int{1, 2, 3, 4, 5}

	ans := [][]int{
		{2, 3, 4, 5, 1},
		{4, 5, 1, 2, 3},
		{1, 2, 3, 4, 5},
	}

	var ansindex int
	for i := 1; i <= 5; i++ {
		Rotate(inp)

		if (i % 2) == 1 {
			if diff := cmp.Diff(inp, ans[ansindex]); diff != "" {
				t.Fatalf("invalid rotattion, diff: %v", diff)
			}
			ansindex++
		}
	}
}
