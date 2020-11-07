package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.String("port", "8080", "listening port")
	tz := flag.String("timezone", "Etc/GMT", "expect timezone")

	flag.Parse()

	err := changeTZ(*tz)
	if err != nil {
		log.Fatal(err)
	}

	localPort := fmt.Sprintf("localhost:%v", *port)

	listener, err := net.Listen("tcp", localPort)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func changeTZ(tz string) error {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}
	time.Local = loc
	log.Println(loc)
	return nil
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // example: disconnect from client
		}
		time.Sleep(1 * time.Second)
	}
}
