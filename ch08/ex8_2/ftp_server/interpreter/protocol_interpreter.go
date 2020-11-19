package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
	prevCmd  int
	hostPort string
}

type fileInfo struct {
	info os.FileInfo
	path string
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

func (pi *interpreter) checkPathExist(path string) (*fileInfo, error) {
	var dstPath string
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dstPath = filepath.Join(pi.wd, path)
	} else if !os.IsNotExist(err) {
		dstPath = path
	} else {
		return nil, err
	}

	fi, err := os.Stat(dstPath)
	return &fileInfo{fi, dstPath}, err
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
			if len(args) == 0 {
				statusCode, err = pi.list(".")
			} else if len(args) != 1 {
				_, err = pi.printf("invalid arguments, ls commands must be \"ls path/to/destination\" ")
				statusCode = StatusBadArguments
			} else {
				statusCode, err = pi.list(args[0])
			}
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

		pi.prevCmd = cmdInt
	}

	if scanner.Err() != nil {
		pi.conn.Write([]byte(scanner.Err().Error()))
	}

	pi.conn.Close()
}

func (pi *interpreter) dataConnection() (io.ReadWriteCloser, error) {
	var conn net.Conn
	var err error

	log.Println(pi.hostPort)

	switch pi.prevCmd {
	case command.PORT:
		conn, err = net.Dial("tcp", pi.hostPort)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("previous command not PORT")
	}

	return conn, nil
}

func (pi *interpreter) changeDir(dst string) (int, error) {
	var err error
	fi := new(fileInfo)
	if fi, err = pi.checkPathExist(dst); os.IsNotExist(err) {
		pi.printf("%v: No such file or directory ", dst)
		return StatusBadArguments, nil
	} else if err != nil {
		log.Println(err)
		pi.printf("%v: server error ", dst)
		return StatusFileUnavailable, nil
	}

	pi.wd = fi.path
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

func (pi *interpreter) list(dst string) (int, error) {
	conn, err := pi.dataConnection()
	if err != nil {
		return StatusBadCommand, err
	}
	defer conn.Close()

	fi := new(fileInfo)

	if fi, err = pi.checkPathExist(dst); os.IsNotExist(err) {
		pi.printf("%v: No such file or directory ", dst)
		return StatusBadArguments, nil
	} else if err != nil {
		log.Println(err)
		pi.printf("%v: server error ", dst)
		return StatusFileUnavailable, nil
	}

	if fi.info.IsDir() {
		files, err := ioutil.ReadDir(fi.path)
		if err != nil {
			pi.printf("%v: No such file or directory ", dst)
			return StatusBadArguments, nil
		}

		pi.printf("%v\r\n", StatusAboutToSend)

		for _, f := range files {
			_, err := fmt.Fprint(conn, f.Name(), "\n")
			if err != nil {
				return StatusBadCommand, err
			}
		}
	} else {
		pi.printf("%v\r\n", StatusAboutToSend)
		_, err := fmt.Fprint(conn, fi.info.Name(), "\n")
		if err != nil {
			return StatusBadCommand, err
		}
	}

	return StatusClosingDataConnection, nil
}

func (pi *interpreter) get() {
	panic("not impl")
}

func (pi *interpreter) close() {
	panic("not impl")
}
