package command

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.12/comic"
)

func GetComic(n string) (*comic.Comic, error) {
	u, err := url.Parse(comic.URL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, n, "info.0.json")

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var c comic.Comic
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
