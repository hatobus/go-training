package interpreter

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/hatobus/go-training/ch08/ex8_2/ftp_server/command"
)

type interpreter struct {
	conn net.Conn // PI (Protocol interpreter) connection
	wd   string
}

func NewInterpreter(newConnection net.Conn) *interpreter {
	return &interpreter{
		conn: newConnection,
	}
}

func (pi *interpreter) println(format string, args ...interface{}) (int, error) {
	line := fmt.Sprintf(format, args...)
	return pi.conn.Write([]byte(line))
}

func (pi *interpreter) setWorkingDir() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	pi.wd = u.HomeDir
	return nil
}

// Start to wait user input command
func (pi *interpreter) Run() {
	pi.conn.Write(StatusTextln(StatusCommandOK))

	err := pi.setWorkingDir()
	if err != nil {
		pi.conn.Write([]byte(err.Error()))
		pi.conn.Close()
		return
	}

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
			if _, err = pi.println("command \"%v\" is not expected! \"help\" command show the usage commands\r\n", cmd); err != nil {
				log.Println(err)
			}
			continue
		}

		var statusCode int

		switch cmdInt {
		case command.CWD:
			if len(args) != 1 {
				_, err = pi.println("invalid arguments, cd commands must be \"cd path/to/destination\"\r\n")
				statusCode = StatusBadArguments
			} else {
				statusCode, err = pi.changeDir(args[0])
			}
		case command.DELE:
			statusCode = StatusNotImplemented
		case command.HELP:
			statusCode = StatusHelp
		case command.LIST:
			statusCode = StatusNotImplemented
		case command.PWD:
			statusCode, err = pi.printWorkingDir()
		case command.RETR:
			statusCode = StatusNotImplemented
		case command.USER, command.PASS, command.ACCT:
			// 今回ログインは実装しない
			statusCode = StatusLoggedIn
		case command.PORT:
			statusCode = StatusCommandOK
		case command.QUIT:
			statusCode = StatusClosing
			break
		default:
			if _, err = pi.println("command \"%v\": [%v] is not expected! \"help\" command show the usage commands\r\n", cmd, args); err != nil {
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

func (pi *interpreter) changeDir(dst string) (int, error) {
	dstPath := filepath.Join(pi.wd, dst)
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		pi.println("%v: No such file or directory ", dst)
		return StatusBadArguments, nil
	}
	pi.wd = dstPath
	return StatusRequestedFileActionOK, nil
}

func (pi *interpreter) printWorkingDir() (int, error) {
	_, err := pi.println(pi.wd + " ")
	if err != nil {
		return StatusBadCommand, err
	}
	return StatusRequestedFileActionOK, nil
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
