package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
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

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	t := []uint64{}

	for _, word := range s.words {
		if word != uint64(x) {
			t = append(t, word)
		}
	}

	s.words = t
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{
		words: s.words,
	}
}

func (s *IntSet) IntersectWith(is *IntSet) {
	for i, word := range is.words {
		if i < len(s.words) {
			s.words[i] &= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) DifferenceWith(is *IntSet) {
	for i, word := range is.words {
		if i < len(s.words) {
			s.words[i] &^= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) SymmetricDifference(is *IntSet) {
	for i, word := range is.words {
		if i < len(s.words) {
			s.words[i] ^= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) Elems() []int {
	e := []int{}
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, i*64+j)
			}
		}
	}
	return e
}
