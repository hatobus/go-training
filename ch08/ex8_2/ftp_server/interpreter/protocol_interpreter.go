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
	conn     net.Conn // PI (Protocol interpreter) connection
	wd       string
	hostPort string
}

func NewInterpreter(newConnection net.Conn) *interpreter {
	return &interpreter{
		conn: newConnection,
	}
}

func (pi *interpreter) printf(format string, args ...interface{}) (int, error) {
	line := fmt.Sprintf(format, args...)
	//return pi.conn.Write([]byte(line))
	return fmt.Fprint(pi.conn, line)
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

		log.Printf("cmd: %v, args; %v\n", cmd, args)

		var cmdInt int
		var err error
		cmdInt, ok := command.CMD[cmd]
		if !ok {
			if _, err = pi.printf("command \"%v\" is not expected! \"help\" command show the usage commands ", cmd); err != nil {
				log.Println(err)
			}
			continue
		}

		var statusCode int

		switch cmdInt {
		case command.CWD:
			if len(args) != 1 {
				_, err = pi.printf("invalid arguments, cd commands must be \"cd path/to/destination\" ")
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
		case command.SYST:
			statusCode = StatusName
		case command.PORT:
			if len(args) != 1 {
				_, err = pi.printf("invalid arguments \"PORT\" commands needs address arguments ")
				statusCode = StatusBadArguments
			} else {
				statusCode = pi.port(args[0])
			}
		case command.LPRT:
			continue
		case command.QUIT:
			statusCode = StatusClosing
			break
		default:
			if _, err = pi.printf("command \"%v\": [%v] is not expected! \"help\" command show the usage commands ", cmd, args); err != nil {
				log.Println(err)
			}
			statusCode = StatusHelp
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
		pi.printf("%v: No such file or directory ", dst)
		return StatusBadArguments, nil
	}
	pi.wd = dstPath
	return StatusRequestedFileActionOK, nil
}

func (pi *interpreter) printWorkingDir() (int, error) {
	_, err := pi.printf(pi.wd + " ")
	if err != nil {
		return StatusBadCommand, err
	}
	return StatusRequestedFileActionOK, nil
}

func (pi *interpreter) port(address string) int {
	var err error
	pi.hostPort, err = pi.hostPortFTP(address)
	if err != nil {
		pi.printf("parse address failed ")
		return StatusBadArguments
	}
	return StatusCommandOK
}

func (pi *interpreter) hostPortFTP(address string) (string, error) {
	var h1, h2, h3, h4 byte
	var p1, p2 int

	_, err := fmt.Sscanf(address, "%d,%d,%d,%d,%d,%d", &h1, &h2, &h3, &h4, &p1, &p2)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d.%d.%d.%d:%d", h1, h2, h3, h4, 256*p1+p2), nil
}

//func (pi *interpreter) nlst(dst string) (int, error) {
//	dstPath := filepath.Join(pi.wd, dst)
//
//	fi, err := ioutil.ReadDir(dstPath)
//	if err != nil {
//		pi.printf("%v: No such file or directory ", dst)
//		return StatusBadArguments, nil
//	}
//
//
//}

func (pi *interpreter) get() {
	panic("not impl")
}

func (pi *interpreter) close() {
	panic("not impl")
}
