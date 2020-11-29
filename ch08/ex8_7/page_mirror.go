package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}
var base *url.URL

// copy from Ex5-12
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

func getLinkNodes(node *html.Node) []*html.Node {
	var links []*html.Node
	nodesVisit := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, n)
		}
	}
	forEachNode(node, nodesVisit, nil)
	return links
}

func getLinkURLs(ln []*html.Node, base *url.URL) []string {
	var urls []string
	for _, n := range ln {
		for _, attr := range n.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := base.Parse(attr.Val)
			if err != nil {
				log.Printf("error occured from %v: %v\n", attr.Val, err)
				continue
			}
			if link.Host != base.Host {
				// Hostが違う場合には保存しない
				continue
			}
			urls = append(urls, link.String())
		}
	}
	return urls
}

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	tokens <- struct{}{} // tokenを取得する
	urls, err := visit(url)
	<-tokens // release token
	if err != nil {
		log.Printf("visit %s: %s", url, err)
	}

	if depth >= maxDepth {
		return
	}

	for _, link := range urls {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

// 相対リンクに書き換えて、拡張子がないリンクはindex.htmlに名前を変える
func toRelative(ln []*html.Node, base *url.URL) {
	for _, n := range ln {
		for i, attr := range n.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := base.Parse(attr.Val)
			if err != nil || link.Host != base.Host {
				continue
			}

			link.Host = ""
			link.Scheme = ""

			attr.Val = link.String()
			n.Attr[i] = attr
		}
	}
}

func save(resp *http.Response, body io.Reader) error {
	u := resp.Request.URL
	fname := filepath.Join(u.Host, u.Path)
	if len(filepath.Ext(u.Path)) == 0 {
		fname = filepath.Join(u.Host, u.Path, "index.html")
	}
	err := os.MkdirAll(filepath.Dir(fname), 0777)
	if err != nil {
		return err
	}

	f, err := os.Create(fname)
	if err != nil {
		return err
	}

	if body != nil {
		_, err = io.Copy(f, body)
	} else {
		_, err = io.Copy(f, resp.Body)
	}

	if err != nil {
		log.Printf("error occured at saving: %v\n", err)
	}
	return nil
}

func visit(urlstr string) ([]string, error) {
	resp, err := http.Get(urlstr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %v from %v", http.StatusText(resp.StatusCode), urlstr)
	}

	u, err := base.Parse(urlstr)
	if err != nil {
		return nil, err
	}

	if base.Host != u.Host {
		// Hostが違う場合には保存しない
		return nil, nil
	}

	var body io.Reader
	var linkURLs []string
	contentType := resp.Header["Content-Type"]
	if strings.Contains(strings.Join(contentType, ","), "text/html") {
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("parsing error from %v: %v", u, err)
		}

		linkNodes := getLinkNodes(doc)
		linkURLs = append(linkURLs, getLinkURLs(linkNodes, u)...)

		toRelative(linkNodes, u)

		b := &bytes.Buffer{}
		err = html.Render(b, doc)
		if err != nil {
			log.Printf("rendering error %v: %v")
		}
		body = b
	}
	return linkURLs, save(resp, body)
}

func main() {
	flag.IntVar(&maxDepth, "d", 3, "max crawl depth")
	flag.Parse()

	wg := &sync.WaitGroup{}
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "usage: mirror URL ...")
		os.Exit(1)
	}

	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid url: %s\n", err)
	}

	base = u
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 1, wg)
	}
	wg.Wait()
}
