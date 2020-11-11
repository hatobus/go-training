package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/hatobus/go-training/ch08/ex8_2/ftp_server/interpreter"
)

func main() {
	var ftpPort int
	// Macだと21番ポートなど1000番以下のポートを指定すると怒られるので注意する
	// Linuxでも怒られたので21番で待ち受けたいときにはsudoを付ける必要がある
	flag.IntVar(&ftpPort, "port", 21, "ftp listen port")

	flag.Parse()

	listener, err := net.Listen("tcp4", fmt.Sprintf(":%v", ftpPort))
	if err != nil {
		log.Fatalf("Port %v failed to open, error: %v\n", ftpPort, err)
	}

	// Connectionを待ち続ける
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Accept failed: %v\n", err)
		}
		go interpreter.NewInterpreter(conn).Run()
	}
}
