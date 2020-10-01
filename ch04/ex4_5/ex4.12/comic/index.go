package comic

func GenerateIndex(comics chan Comic) IndexOfNumber {
	in := make(IndexOfNumber)

	for c := range comics {
		in[c.Num] = c
	}
	return in
}
