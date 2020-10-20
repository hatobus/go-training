package ex5_7

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var tagDepth int

type HTMLPrettyPrinter struct {
	w   io.Writer
	err error
}

func NewHTMLPrettyPrinter() HTMLPrettyPrinter {
	return HTMLPrettyPrinter{}
}

func (pp HTMLPrettyPrinter) PPrint(w io.Writer, fileName string) error {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return fmt.Errorf("%v is not found", fileName)
	} else if err != nil {
		return err
	}

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	node, err := html.Parse(f)
	if err != nil {
		return err
	}

	pp.w = w
	pp.err = nil

	pp.forEachNode(node)

	return pp.err
}

func (pp HTMLPrettyPrinter) forEachNode(node *html.Node) {

	pp.start(node)

	if pp.err != nil {
		return
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		pp.forEachNode(c)
	}

	pp.end(node)

	if pp.err != nil {
		return
	}
}

func (pp HTMLPrettyPrinter) start(node *html.Node) {
	switch node.Type {
	case html.ElementNode:
		pp.startElem(node)
	case html.TextNode:
		pp.startText(node)
	case html.CommentNode:
		pp.startComment(node)
	}
}

func (pp HTMLPrettyPrinter) startElem(node *html.Node) {
	end := ">"
	if node.FirstChild == nil {
		end = "/>"
	}

	attributes := make([]string, 0, len(node.Attr))
	for _, attr := range node.Attr {
		attributes = append(attributes, fmt.Sprintf("%s=\"%s\"", attr.Key, attr.Val))
	}

	attrStr := ""
	if len(node.Attr) > 0 {
		attrStr = " " + strings.Join(attributes, " ")
	}

	name := node.Data

	pp.printf("%*s<%s%s%s\n", tagDepth*2, "", name, attrStr, end)
	tagDepth++
}

func (pp HTMLPrettyPrinter) startText(node *html.Node) {
	text := strings.TrimSpace(node.Data)
	if len(text) == 0 {
		return
	}
	pp.printf("%*s%s\n", tagDepth*2, "", node.Data)
}

func (pp HTMLPrettyPrinter) startComment(node *html.Node) {
	pp.printf("<!--%v-->\n", node.Data)
}

func (pp HTMLPrettyPrinter) end(node *html.Node) {
	switch node.Type {
	case html.ElementNode:
		pp.endElem(node)
	default:
		return
	}
}

func (pp HTMLPrettyPrinter) endElem(node *html.Node) {
	tagDepth--
	if node.FirstChild == nil {
		return
	}
	pp.printf("%*s</%s>\n", tagDepth*2, "", node.Data)
}

func (pp HTMLPrettyPrinter) printf(format string, args ...interface{}) {
	_, pp.err = fmt.Fprintf(pp.w, format, args...)
}
