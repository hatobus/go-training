package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestDup2(t *testing.T) {
	type T struct{}

	type testData struct {
		Args []string
		expectOutput map[string]T
	}

	testCases := map[string]testData {
		"重複した行が1行ある場合": {
			Args: []string{"boy.txt"},
			expectOutput: map[string]T{
				fmt.Sprintf("2\tboy.txt: 新築の二階から首を出していたら、同級生の一人が冗談に、いくら威張っても、そこから飛び降りる事は出来まい。弱虫やーい。と囃したからである。"): T{},
			},
		},
		"重複した行が複数行ある場合": {
			Args: []string{"heart.txt"},
			expectOutput: map[string]T{
				fmt.Sprintf("3\theart.txt: 私はその人の記憶を呼び起すごとに、すぐ「先生」といいたくなる。"): T{},
				fmt.Sprintf("2\theart.txt: 私はその人を常に先生と呼んでいた。"): T{},
			},
		},
		"複数ファイルで重複した行が複数行ある場合": {
			Args: []string{"boy.txt", "heart.txt"},
			expectOutput: map[string]T{
				fmt.Sprintf("3\theart.txt: 私はその人の記憶を呼び起すごとに、すぐ「先生」といいたくなる。"): T{},
				fmt.Sprintf("2\tboy.txt: 新築の二階から首を出していたら、同級生の一人が冗談に、いくら威張っても、そこから飛び降りる事は出来まい。弱虫やーい。と囃したからである。"): T{},
				fmt.Sprintf("2\theart.txt: 私はその人を常に先生と呼んでいた。"): T{},
			},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T){

			args := make([]string, 0, len(tc.Args)+2)
			args = append(args, "run", "dup2.go")
			args = append(args, tc.Args...)

			out, err := exec.Command("go", args...).Output()
			if err != nil {
				t.Fatalf("execute command failed: %v", err)
			}

			duplicate := strings.Split(string(out), "\n")
			for _, elem := range duplicate[:len(duplicate)-1] {
				if _, ok := tc.expectOutput[elem]; !ok {
						t.Fatalf("unexpected output: %v", elem)
				}
			}
		})
	}
}

