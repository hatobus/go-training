package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(node *html.Node, tags ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	keep := make(map[string]bool, len(tags))
	for _, t := range tags {
		keep[t] = true
	}

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		_, ok := keep[n.Data]
		if ok {
			nodes = append(nodes, n)
		}
		return true
	}
	forEachElement(node, pre, nil)
	return nodes
}

func forEachElement(node *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0) // unvisited
	u = append(u, node)
	for len(u) > 0 {
		node = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(node) {
				return node
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
		if post != nil {
			if !post(node) {
				return node
			}
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "please input HTML_FILE TAG ...")
	}
	filename := os.Args[1]
	tags := os.Args[2:]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range ElementsByTagName(doc, tags...) {
		fmt.Printf("%+v\n", n)
	}
}
