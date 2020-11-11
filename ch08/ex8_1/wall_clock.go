package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Clock struct {
	Name    string
	Address string
}

func (c *Clock) print(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Fprintf(w, "%v\t%v\n", c.Name, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("read failed from %v, error: %v", c.Name, scanner.Err())
	}
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("invalid argument please input city=address")
	}

	clocks := make([]*Clock, 0, len(os.Args)-1)

	for _, arg := range os.Args[1:] {
		d := strings.SplitN(arg, "=", 2)
		clocks = append(clocks, &Clock{Name: d[0], Address: d[1]})
	}

	for _, clock := range clocks {
		conn, err := net.Dial("tcp", clock.Address)
		if err != nil {
			log.Fatalf("net dial connection error: %v", err)
		}
		defer conn.Close()
		go clock.print(os.Stdout, conn)
	}

	for {
		time.Sleep(10 * time.Second)
	}
}
