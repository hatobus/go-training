package ex5_15

import "fmt"

func Max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("arguments length is 0")
	}

	max := vals[0]
	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}

	return max, nil
}
