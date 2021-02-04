package ex13_2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCycle(t *testing.T) {
	type cycle struct {
		c *cycle
	}

	c1, c2, c3 := cycle{}, cycle{}, cycle{}
	c1.c = &c2
	c2.c = &c3

	uncycleData := c1

	c4, c5, c6 := cycle{}, cycle{}, cycle{}
	c4.c = &c5
	c5.c = &c6
	c6.c = &c4

	cycledData := c4

	testCases := map[string]struct {
		x      interface{}
		expect bool
	}{
		"cycled struct": {
			x:      cycledData,
			expect: true,
		},
		"uncycled struct": {
			x:      uncycleData,
			expect: false,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			tf := Cycle(tc.x)
			if diff := cmp.Diff(tf, tc.expect); diff != "" {
				t.Fatalf("unexpected output diff = %v", diff)
			}
		})
	}
}
