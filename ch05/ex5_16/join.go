package ex5_16

import "fmt"

func Join(sep string, elems ...string) (string, error) {
	if len(elems) == 0 {
		return "", fmt.Errorf("argument length is 0")
	} else if len(elems) == 1 {
		return elems[0], nil
	}

	var s string

	for _, elem := range elems[:len(elems)-1] {
		s += elem + sep
	}

	s += elems[len(elems)-1]

	return s, nil
}
