package comic

func GenerateIndex(comics chan Comic) (IndexOfWord, IndexOfNumber) {
	iw := make(IndexOfWord)
	in := make(IndexOfNumber)

	for c := range comics {
		in[c.Num] = c
	}
	return iw, in
}
