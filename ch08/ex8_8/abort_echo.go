package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var timeOut = 10 * time.Second

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()

	s := bufio.NewScanner(c)
	scanLines := make(chan string)

	go func() {
		for s.Scan() {
			scanLines <- s.Text()
		}
	}()

	for {
		select {
		case scanLine, ok := <-scanLines:
			if !ok {
				return
			}
			go echo(c, scanLine, 1*time.Second)
		case <-time.After(timeOut):
			return
		}
	}
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
