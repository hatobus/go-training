package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const timeout = 10 * time.Second

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

type client struct {
	MsgOut chan<- string
	Name   string
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.MsgOut <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.MsgOut <- "Present:"
			for c := range clients {
				cli.MsgOut <- c.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.MsgOut)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
