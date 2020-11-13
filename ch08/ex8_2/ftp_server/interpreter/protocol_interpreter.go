package interpreter

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

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
	pi.conn.Write([]byte("200 Ready."))
	pi.conn.Write([]byte(fmt.Sprintf("HI! %v\n", time.Now().Format(time.RFC822))))

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

		pi.conn.Write([]byte(fmt.Sprintf("input cmd: %v, input args: %v\n", cmd, args)))

		_, ok := command.CMD[cmd]
		if !ok {
			pi.conn.Write([]byte(fmt.Sprintf("command \"%v\" is not expected! \"help\" command show the usage commands\n", cmd)))
			continue
		}

		switch command.CMD[cmd] {
		case command.CWD:
			pi.conn.Write([]byte(fmt.Sprintf("your command is CWD: %v\n", command.CWD)))
		default:
			pi.conn.Write([]byte(fmt.Sprintf("command \"%v\" is not expected! \"help\" command show the usage commands\n", cmd)))
			continue
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
