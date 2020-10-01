package command

import (
	"strconv"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.12/comic"
)

func GetComic(n string) (*comic.Comic, error) {

	i, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}

	return comic.GetComic(i)
}
