package ex5_4

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func ExtractLinksFromHTML(htmlFileName string) ([]string, error) {
	_, err := os.Stat(htmlFileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("%v is not found", htmlFileName)
	} else if err != nil {
		return nil, err
	}

	f, err := os.Open(htmlFileName)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(f)

	links := visitLinks(nil, doc)

	return links, nil
}

func visitLinks(links []string, node *html.Node) []string {
	if node == nil {
		return nil
	}

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if node.Type == html.ElementNode && (node.Data == "img" || node.Data == "script") {
		for _, a := range node.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}

	if node.Type == html.ElementNode && node.Data == "link" {
		for _, a := range node.Attr {
			if a.Key == "rel" {
				links = append(links, a.Val)
			}
		}
	}

	if node.FirstChild != nil {
		links = visitLinks(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		links = visitLinks(links, node.NextSibling)
	}

	return links
}
