package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	m map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{
		m: map[int]bool{},
	}
}

func (m *MapIntSet) Has(x int) bool {
	return m.m[x]
}

func (m *MapIntSet) Add(x int) {
	m.m[x] = true
}

func (m *MapIntSet) UnionWith(t IntSet) {
	if mis, ok := t.(*MapIntSet); ok {
		for _, x := range ints(mis) {
			m.m[x] = true
		}
	}
}

func (m *MapIntSet) Len() int {
	return len(m.m)
}

func (m *MapIntSet) Remove(x int) {
	delete(m.m, x)
}

func (m *MapIntSet) Clear() {
	m.m = make(map[int]bool)
}

func (m *MapIntSet) Copy() IntSet {
	duplicated := make(map[int]bool)
	for k, v := range m.m {
		duplicated[k] = v
	}
	return &MapIntSet{
		m: duplicated,
	}
}

func (m *MapIntSet) String() string {
	b := &bytes.Buffer{}
	b.WriteString("{")
	for i, x := range ints(m) {
		if i != 0 {
			b.WriteString(" ")
		}
		b.WriteString(fmt.Sprintf("%d", x))
	}
	b.WriteString("}")
	return b.String()
}

func ints(m *MapIntSet) []int {
	ints := make([]int, 0, len(m.m))
	for x := range m.m {
		ints = append(ints, x)
	}
	sort.IntSlice(ints).Sort()
	return ints
}
