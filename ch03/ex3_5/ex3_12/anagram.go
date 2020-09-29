package anagram

import (
	"sort"
	"strings"
)

func Anagram(s1, s2 string) bool {
	s1 = deleteSpace(strings.ToLower(s1))
	s2 = deleteSpace(strings.ToLower(s2))

	if len(s1) != len(s2) {
		return false
	}

	runeS1 := []rune(s1)
	runeS2 := []rune(s2)

	sort.Slice(runeS1, func(i, j int) bool {
		return runeS1[i] < runeS1[j]
	})

	sort.Slice(runeS2, func(i, j int) bool {
		return runeS2[i] < runeS2[j]
	})

	if string(runeS1) == string(runeS2) {
		return true
	}

	return false
}

func deleteSpace(s string) string {
	elems := strings.Split(s, " ")
	return strings.Join(elems, "")
}
