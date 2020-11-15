package interpreter

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/hatobus/go-training/ch08/ex8_2/ftp_server/command"
)

type interpreter struct {
	conn net.Conn // PI (Protocol interpreter) connection
}

func NewInterpreter(newConnection net.Conn) *interpreter {
	return &interpreter{
		conn: newConnection,
	}
}

// Start to wait user input command
func (pi *interpreter) Run() {
	pi.conn.Write(StatusTextln(StatusCommandOK))

	scanner := bufio.NewScanner(pi.conn)

	for scanner.Scan() {
		var cmd string
		var args []string

		userInput := strings.Fields(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		cmd = userInput[0]

		if len(userInput) > 1 {
			args = userInput[1:]
		}

		var cmdInt int
		var err error
		cmdInt, ok := command.CMD[cmd]
		if !ok {
			if _, err = pi.conn.Write([]byte(fmt.Sprintf("command \"%v\" is not expected! \"help\" command show the usage commands\n", cmd))); err != nil {
				log.Println(err)
			}
			continue
		}

		var statusCode int

		switch cmdInt {
		case command.CWD:
			_, err = pi.conn.Write([]byte(fmt.Sprintf("your command is CWD: %v\n", command.CWD)))
		case command.DELE:
			statusCode = StatusRequestedFileActionOK
		case command.HELP:
			statusCode = StatusHelp
		case command.LIST:
			statusCode = StatusNotImplemented
		case command.PWD:
			statusCode = StatusNotImplemented
		case command.RETR:
			statusCode = StatusNotImplemented
		case command.USER, command.PASS, command.ACCT:
			// 今回ログインは実装しない
			statusCode = StatusNotImplemented
		case command.PORT:
			statusCode = StatusCommandOK
		case command.QUIT:
			statusCode = StatusClosing
			break
		default:
			if _, err = pi.conn.Write([]byte(fmt.Sprintf("command \"%v\": [%v] is not expected! \"help\" command show the usage commands\n", cmd, args))); err != nil {
				log.Println(err)
			}
			_, err = pi.conn.Write(StatusTextln(StatusHelp))
			continue
		}

		_, err = pi.conn.Write(StatusTextln(statusCode))
		if err != nil {
			log.Println(err)
		}
	}

	if scanner.Err() != nil {
		pi.conn.Write([]byte(scanner.Err().Error()))
	}

	pi.conn.Close()
}

func (pi *interpreter) changeDir() {
	panic("not impl")
}

func (pi *interpreter) list() {
	panic("not impl")
}

func (pi *interpreter) get() {
	panic("not impl")
}

func (pi *interpreter) close() {
	panic("not impl")
}
