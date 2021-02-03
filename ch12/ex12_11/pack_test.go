package ex12_11

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPacking(t *testing.T) {
	type parameter struct {
		BookName string `http:"name"`
		Pages    int    `http:"pages"`
	}
	testCases := map[string]struct {
		parameters *parameter
		expect     string
	}{
		"1つの要素": {
			parameters: &parameter{
				BookName: "hatobus",
			},
			expect: "name=hatobus&pages=0",
		},
		"すべての要素": {
			parameters: &parameter{
				BookName: "kokoro",
				Pages:    450,
			},
			expect: "name=kokoro&pages=450",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			u, err := Pack(tc.parameters)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(tc.expect, u.RawQuery); diff != "" {
				t.Errorf("unexpected error diff = %v", diff)
			}
		})
	}
}
