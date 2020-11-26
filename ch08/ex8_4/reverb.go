package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var delay = 1 * time.Second

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	// WaitGroupを値渡しするとgoroutineがdeadlockします
	// ポインタ渡しするかクロージャでなんとかする
	for input.Scan() {
		wg.Add(1)
		go func(shout string) {
			fmt.Fprintln(c, "\t", strings.ToUpper(shout))
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", shout)
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", strings.ToLower(shout))
			wg.Done()
		}(input.Text())
	}
	wg.Wait()
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}
