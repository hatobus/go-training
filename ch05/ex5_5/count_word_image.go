package ex5_5

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImage(us string) (words, images int, err error) {
	u, err := url.Parse(us)
	if err != nil {
		return
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = fmt.Errorf("your requested URL returned status code 2XX, returned %v", resp.StatusCode)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	words, images = countWordAndImage(doc)

	return
}

func countWordAndImage(n *html.Node) (words, images int) {
	nodes := make([]*html.Node, 0)
	nodes = append(nodes, n)

	for len(nodes) > 0 {
		n = nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]

		switch n.Type {
		case html.TextNode:
			words += wc(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}

		for fc := n.FirstChild; fc != nil; fc = fc.NextSibling {
			nodes = append(nodes, fc)
		}
	}
	return
}

func wc(s string) (n int) {
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n += 1
	}
	return
}
