package main

import (
	"flag"
	"log"
	"sync"

	"gopl.io/ch5/links"
)

var tokens = make(chan struct{}, 20)
var maxDepth int
var sl = sync.Mutex{}
var seen = make(map[string]bool)

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth >= maxDepth {
		return
	}

	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Println(err)
	}
	for _, link := range list {
		sl.Lock()
		if seen[link] {
			sl.Unlock()
			continue
		}
		seen[link] = true
		sl.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

func main() {
	flag.IntVar(&maxDepth, "depth", 3, "crawling depth max value")
	flag.Parse()

	wg := &sync.WaitGroup{}
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 0, wg)
	}
	wg.Wait()
}
