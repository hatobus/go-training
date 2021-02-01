package ex12_1

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDisdplay(t *testing.T) {
	ptrInt := 10
	testCases := map[string]struct {
		input  reflect.Value
		expect string
	}{
		"input int 3": {
			input:  reflect.ValueOf(int(3)),
			expect: "input = 3\n",
		},
		"input slice": {
			input:  reflect.ValueOf([]string{"a", "b", "c"}),
			expect: "input[0] = \"a\"\ninput[1] = \"b\"\ninput[2] = \"c\"\n",
		},
		"input struct": {
			input: reflect.ValueOf(
				struct {
					name string
					age  int
				}{
					name: "hatobus",
					age:  23,
				},
			),
			expect: "input.name = \"hatobus\"\ninput.age = 23\n",
		},
		"input ptr": {
			input:  reflect.ValueOf(&ptrInt),
			expect: "(*input) = 10\n",
		},
		"input map has string value": {
			input:  reflect.ValueOf(map[string]string{"1": "a", "2": "b", "3": "c"}),
			expect: "input[\"1\"] = \"a\"\ninput[\"2\"] = \"b\"\ninput[\"3\"] = \"c\"\n",
		},
		"input map has string slices value": {
			input: reflect.ValueOf(map[int][]string{
				1: {"a"},
				2: {"a", "b"},
				3: {"a", "b", "c"},
			}),
			expect: "input[1][0] = \"a\"\ninput[2][0] = \"a\"\ninput[2][1] = \"b\"\ninput[3][0] = \"a\"\ninput[3][1] = \"b\"\ninput[3][2] = \"c\"\n",
		},
		"input map has string string slices key": {
			input: reflect.ValueOf(map[[3]string]string{
				{"1", "s", "t"}: "1",
			}),
			expect: "input[[\"1\", \"s\", \"t\"]] = \"1\"\n",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			displayed := Display("input", tc.input)
			if diff := cmp.Diff(tc.expect, displayed); diff != "" {
				t.Errorf("invalid output: diff = %v\n", diff)
			}
		})
	}
}
