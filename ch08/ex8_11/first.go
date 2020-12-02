package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	flag.Parse()

	cancel := make(chan struct{})
	resp := make(chan *http.Response)
	var wg sync.WaitGroup

	for _, url := range flag.Args() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			client := http.DefaultClient
			req, err := http.NewRequest(http.MethodHead, url, nil)
			if err != nil {
				log.Printf("HEAD error %s from: %s", url, err)
				return
			}
			res, err := client.Do(req)
			if err != nil {
				log.Printf("HEAD error %s from: %s", url, err)
				return
			}
			resp <- res
		}(url)
	}

	res := <-resp

	defer res.Body.Close()
	close(cancel) // 不完全なHTTPリクエストはキャンセルする

	for name, v := range res.Header {
		fmt.Printf("%s: %s\n", name, strings.Join(v, ","))
	}

	wg.Wait()
}
