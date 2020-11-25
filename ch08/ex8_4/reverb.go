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

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func wgCloser(wg sync.WaitGroup, s chan<- struct{}) {
	wg.Wait()
	close(s)
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	finished := make(chan struct{})
	var wg sync.WaitGroup

	for {
		log.Println("for")
		wg.Add(1)
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
		go wgCloser(wg, finished)
		log.Println("end")
	}
}
