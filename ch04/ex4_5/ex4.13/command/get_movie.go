package command

import (
	"github.com/hatobus/go-training/ch04/ex4_5/ex4.13/movie"
	"golang.org/x/xerrors"
)

func GetMovie(title string) (string, error) {
	m, err := movie.GetMovie(title)
	if err != nil {
		return "", err
	} else if m.Title == "" {
		return "", xerrors.Errorf("%v is not found", title)
	}

	fname, err := m.GetPoster()
	if err != nil {
		return "", err
	}

	return fname, nil
}
