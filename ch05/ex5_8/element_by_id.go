package ex5_8

import (
	"golang.org/x/net/html"
)

func ElementByID(node *html.Node, id string) *html.Node {
	pre := func(node *html.Node) bool {
		if node.Type != html.ElementNode {
			return true
		}
		for _, attribute := range node.Attr {
			if attribute.Key == "id" && attribute.Val == id {
				return false
			}
		}
		return true
	}
	return forEachNode(node, pre, nil)
}

func forEachNode(node *html.Node, pre, post func(*html.Node) bool) *html.Node {
	unvisited := make([]*html.Node, 0)
	unvisited = append(unvisited, node)

	for len(unvisited) > 0 {
		node = unvisited[0]
		unvisited = unvisited[1:]

		if pre != nil {
			if !pre(node) {
				return node
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			unvisited = append(unvisited, c)
		}

		if post != nil {
			if !post(node) {
				return node
			}
		}
	}
	return nil
}
