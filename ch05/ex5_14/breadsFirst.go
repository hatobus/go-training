package main

import (
	"fmt"
	"math/rand"
	"time"
)

var prereqs = map[string][]string{
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
}

func breadthFirst(f func(string) []string, worklist []string) []string {
	s := []string{}
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				dep := f(item)
				s = append(s, dep...)
				worklist = append(worklist, dep...)
			}
		}
	}
	return s
}

func dependening(subj string) []string {
	fmt.Println(subj)
	return prereqs[subj]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var subj string
	for k, _ := range prereqs {
		subj = k
		break
	}

	fmt.Println(breadthFirst(dependening, []string{subj}))
}
