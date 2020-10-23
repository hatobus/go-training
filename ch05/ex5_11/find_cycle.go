package main

import (
	"fmt"
	"strings"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
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
}

func main() {
	course, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
	}

	for i, name := range course {
		fmt.Printf("%v:\t%v\n", i, name)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	var err error
	resolvedSubjects := make(map[string]bool)
	var visitAll func([]string, []string)

	index := func(s string, ss []string) (int, error) {
		for i, v := range ss {
			if s == v {
				return i, nil
			}
		}
		return -1, fmt.Errorf("not found")
	}

	visitAll = func(items, parents []string) {
		for _, item := range items {
			v, seen := resolvedSubjects[item]
			if seen && !v {
				start, _ := index(item, parents)
				err = fmt.Errorf("cyclyed reference: %s", strings.Join(append(parents[start:], item), "--> "))
			}
			if !seen {
				resolvedSubjects[item] = false
				visitAll(m[item], append(parents, item))
				resolvedSubjects[item] = true
				order = append(order, item)
			}
		}
	}

	for key := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{key}, nil)
	}

	return order, nil
}
