package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"gopl.io/ch5/links"
)

var hostName string

var crawlFunc func(string) ([]string, error) = nil

func breadthFirst(worklist []string) error {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			fmt.Println(item)
			if !seen[item] {
				seen[item] = true
				work, err := crawlFunc(item)
				if err != nil {
					return err
				}
				worklist = append(worklist, work...)
			}
		}
	}

	return nil
}

func crawl(url string) ([]string, error) {
	err := download(url)
	if err != nil {
		if err.Error() == "different hostname" {
			return nil, nil
		}
		return nil, err
	}

	list, err := links.Extract(url)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func download(rawurl string) error {
	u, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	if hostName == "" {
		hostName = u.Host
	}

	// hostnameが違うサイトはダウンロードしない
	if hostName != u.Host {
		return fmt.Errorf("different hostname")
	}

	var fname string

	dir := filepath.Join("replica_html", hostName)

	if filepath.Ext(fname) == "" {
		dir = filepath.Join(dir, u.Path)
		fname = filepath.Join(dir, u.Path)
	} else {
		dir = filepath.Join(dir, filepath.Dir(u.Path))
		fname = u.Path
	}

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func BreadthFirst(f func(string) ([]string, error), s []string) error {
	switch f {
	case nil:
		crawlFunc = crawl
	default:
		crawlFunc = f
	}

	return breadthFirst(s)
}

func main() {
	err := BreadthFirst(nil, os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
