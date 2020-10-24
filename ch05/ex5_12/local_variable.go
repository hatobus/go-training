package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		line, err := outline(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(line)
	}
}

func outline(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var out string
	var depth int

	startElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			out += fmt.Sprintf("%*s</%s>\n", depth*2, "", node.Data)
			depth++
		}
	}

	endElement := func(node *html.Node) {
		if node.Type == html.ElementNode {
			depth--
			out += fmt.Sprintf("%*s</%s>\n", depth*2, "", node.Data)
		}
	}

	forEachNode(doc, startElement, endElement)

	return out, nil
}

func forEachNode(node *html.Node, startFunc, endFunc func(*html.Node)) {
	if startFunc != nil {
		startFunc(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, startFunc, endFunc)
	}

	if endFunc != nil {
		endFunc(node)
	}

}
