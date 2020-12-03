package popcount

import "sync"

var once sync.Once
var pc [256]byte

var calculate = func() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&i)
	}
}

func PopCount(x uint64) int {
	once.Do(calculate)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountOrigin(x uint64) int {
	once.Do(calculate)
	var cnt int

	for i := 0; i < 8; i++ {
		cnt += int(pc[byte(x>>(uint(i)*8))])
	}

	return cnt
}
