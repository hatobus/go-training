package ex2_5

func PopCount(x uint64) int {
	var cnt int

	for x > 0 {
		x &= uint64(x - uint64(1))
		cnt++
	}

	return cnt
}
