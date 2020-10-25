package ex5_15

import "fmt"

func Min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("arguments length is 0")
	}

	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}

	return min, nil
}
