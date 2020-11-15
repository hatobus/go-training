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
	pi.conn.Write([]byte("200 Ready.\n"))

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

		switch cmdInt {
		case command.CWD:
			_, err = pi.conn.Write([]byte(fmt.Sprintf("your command is CWD: %v\n", command.CWD)))
		case command.USER, command.PASS, command.ACCT:
			// 今回ログインは実装しない
			_, err = pi.conn.Write(StatusTextln(StatusCommandNotImplemented))
		case command.PORT:
			_, err = pi.conn.Write(StatusTextln(StatusCommandOK))
		case command.QUIT:
			_, err = pi.conn.Write(StatusTextln(StatusClosing))
		default:
			if _, err = pi.conn.Write([]byte(fmt.Sprintf("command \"%v\": [%v] is not expected! \"help\" command show the usage commands\n", cmd, args))); err != nil {
				log.Println(err)
			}
			_, err = pi.conn.Write(StatusTextln(StatusHelp))
			continue
		}

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
