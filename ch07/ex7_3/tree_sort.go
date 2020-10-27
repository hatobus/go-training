package ex7_3

import (
	"bytes"
	"fmt"
)

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func (t *Tree) String() string {
	var o []int
	o = appendValues(o, t)
	if len(o) == 0 {
		return ""
	}

	b := &bytes.Buffer{}
	fmt.Fprintf(b, "[%v", o[0])

	for _, v := range o[1:] {
		fmt.Fprintf(b, " %v", v)
	}
	fmt.Fprint(b, "]")

	return b.String()
}

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}
