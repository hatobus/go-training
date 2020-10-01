package movie

import (
	"bufio"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func GetMovie(title string) (*Movie, error) {
	u, err := url.Parse(urlraw)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("apikey", os.Getenv("OMD_APIKEY"))
	q.Set("t", title)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}
	return &movie, err
}

func (m *Movie) GetPoster() (string, error) {
	u, err := url.Parse(m.Poster)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ext := filepath.Ext(u.String())

	f, err := os.Create(m.Title + ext)
	if err != nil {
		return "", err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	// Flushすると一度に書き込みができる
	// https://golang.org/pkg/bufio/#Writer.Flush
	err = w.Flush()
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}
