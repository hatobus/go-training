package main

import "testing"

func TestFindCycle(t *testing.T) {
	testData := map[string]struct {
		input     map[string][]string
		errWant   bool
		errString string
	}{
		"cycleが無い場合": {
			input: map[string][]string{
				"algorithms": {"data structures"},
				"calculus":   {"linear algebra"},

				"compilers": {
					"data structures",
					"formal languages",
					"computer organization",
				},

				"data structures":       {"discrete math"},
				"databases":             {"data structures"},
				"discrete math":         {"intro to programming"},
				"formal languages":      {"discrete math"},
				"networks":              {"operating systems"},
				"operating systems":     {"data structures", "computer organization"},
				"programming languages": {"data structures", "computer organization"},
			},
		},
		"calculusとlinear algebraが循環している": {
			input: map[string][]string{
				"algorithms":     {"data structures"},
				"calculus":       {"linear algebra"},
				"linear algebra": {"calculus"},

				"compilers": {
					"data structures",
					"formal languages",
					"computer organization",
				},

				"data structures":       {"discrete math"},
				"databases":             {"data structures"},
				"discrete math":         {"intro to programming"},
				"formal languages":      {"discrete math"},
				"networks":              {"operating systems"},
				"operating systems":     {"data structures", "computer organization"},
				"programming languages": {"data structures", "computer organization"},
			},
			errWant:   true,
			errString: "cyclyed reference: calculus--> linear algebra--> calculus",
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			_, err := topoSort(tc.input)
			if err != nil {
				switch tc.errWant {
				case false:
					t.Fatalf("err not expected but got, err: %v\n", err)
				default:
					if err.Error() != tc.errString {
						t.Fatalf("invalid error string, want %v but got %v\n", tc.errString, err.Error())
					}
				}
			}
		})
	}
}
