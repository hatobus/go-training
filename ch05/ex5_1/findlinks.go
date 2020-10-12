package ex5_1

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func FindLInks(htmlFname string) ([]string, error) {
	_, err := os.Stat(htmlFname)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("%v is not found", htmlFname)
	} else if err != nil {
		return nil, err
	}

	f, err := os.Open(htmlFname)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)

	for _, link := range visit(nil, doc) {
		links = append(links, link)
	}

	return links, nil
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}
