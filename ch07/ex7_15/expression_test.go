package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var outs = []string{
	"Please input:",
	"<var>=<value>, (ex: a=1)",
}

func TestExpression(t *testing.T) {
	testCases := map[string]struct {
		expression string
		variables  map[string]int
		expect     string
	}{
		"x + yの時": {
			expression: "x+y",
			variables: map[string]int{
				"x": 1, "y": 2,
			},
			expect: fmt.Sprintf("%v %v 3\n", outs[0], outs[1]),
		},
		"log2(8) - log10(1000)": {
			expression: "log2(x) - log10(y)",
			variables: map[string]int{
				"x": 8, "y": 1000,
			},
			expect: fmt.Sprintf("%v %v 0\n", outs[0], outs[1]),
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			cmd := exec.Command("go", "run", "main.go")

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			stdin, err := cmd.StdinPipe()
			if err != nil {
				t.Fatal(err)
			}

			var varsInput string
			for name, val := range tc.variables {
				varsInput += fmt.Sprintf("%v=%v ", name, val)
			}

			vars := []string{tc.expression, varsInput}

			go func() {
				defer stdin.Close()
				for _, in := range vars {
					io.WriteString(stdin, in+"\n")
				}
			}()

			if err := cmd.Run(); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(stdout.String(), tc.expect); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
