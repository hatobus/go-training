package unique

func UniqueStrings(s []string) []string {
	wi := 0
	for _, c := range s {
		if s[wi] == c {
			continue
		}
		wi++
		s[wi] = c
	}
	return s[:wi+1]
}
