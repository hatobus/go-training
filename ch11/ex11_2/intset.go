package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An WordIntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type WordIntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *WordIntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *WordIntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *WordIntSet) UnionWith(t IntSet) {
	if word, ok := t.(*WordIntSet); ok {
		for i, tword := range word.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *WordIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *WordIntSet) Len() int {
	return len(s.words)
}

func (s *WordIntSet) Remove(x int) {
	t := []uint64{}

	for _, word := range s.words {
		if word != uint64(x) {
			t = append(t, word)
		}
	}

	s.words = t
}

func (s *WordIntSet) Clear() {
	s.words = []uint64{}
}

func (s *WordIntSet) Copy() IntSet {
	return &WordIntSet{
		words: s.words,
	}
}
