package ex2_3

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var cnt int

	for i := 0; i < 8; i++ {
		cnt += int(pc[byte(x>>(uint(i)*8))])
	}

	return cnt
}
