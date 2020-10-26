package ex5_18

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Fetch(rawurl string, dirname string) (string, int64, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", 0, err
	}

	local := path.Base(u.String())
	if !strings.HasSuffix(local, ".html") {
		local = "/index.html"
		u.Path = path.Join(u.Path, local)
	}

	if dirname != "" {
		local = filepath.Join(dirname, local)
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode > 299 {
		return "", 0, fmt.Errorf("resposne status code %v", resp.StatusCode)
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	n, err := io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, err
	}

	return local, n, nil

}
