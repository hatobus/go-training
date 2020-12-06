package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/hatobus/go-training/ch08/ex8_2/ftp_server/command"
)

type interpreter struct {
	conn         net.Conn // PI (Protocol interpreter) connection
	wd           string
	prevCmd      int
	listener     net.Listener
	binaryOption bool
	hostPort     string
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
	dstPath := filepath.Join(pi.wd, path)
	log.Println(dstPath)
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		dstPath = path
	} else if !os.IsNotExist(err) {
	} else {
		return nil, err
	}

	fi, err := os.Stat(dstPath)
	return &fileInfo{fi, dstPath}, err
}

func (pi *interpreter) isCorrectlyPath(filePath string) (string, bool) {
	dir := filepath.Dir(filePath)
	if dir == "." {
		return filepath.Join(pi.wd, filePath), true
	}

	fname := filepath.Base(filePath)

	// check rel path
	if _, err := os.Stat(filepath.Join(pi.wd, dir)); !os.IsNotExist(err) {
		abs, _ := filepath.Abs(filepath.Join(pi.wd, dir))
		return filepath.Join(abs, fname), true
	}

	// check abs path
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return filepath.Join(dir, fname), true
	}

	return "", false
}

func (pi *interpreter) dataType(argument []string) (int, error) {
	switch strings.ToUpper(strings.Join(argument, " ")) {
	case "A", "A N":
		pi.binaryOption = false
	case "I", "L 8":
		pi.binaryOption = true
	default:
		return StatusNotImplementedParameter, fmt.Errorf("unsupported data type. Supported types is \"A, A N, I, L 8\"")
	}
	return StatusCommandOK, nil
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
		case command.RMD, command.XRMD:
			if len(args) != 1 {
				_, err = pi.printf("invalid arguments, rm commands must be \"rm path/to/destination\" ")
				statusCode = StatusBadArguments
			} else {
				statusCode, err = pi.delete(args[0])
			}
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
			if len(args) != 1 {
				_, err = pi.printf("invalid arguments \"PORT\" commands needs address arguments ")
				statusCode = StatusBadArguments
			} else {
				statusCode, err = pi.retr(args[0])
			}
		case command.STOR:
			switch len(args) {
			case 1:
				statusCode, err = pi.store(args[0], pi.wd)
			case 2:
				statusCode, err = pi.store(args[0], args[1])
			default:
				_, err = pi.printf("invalid argument \"STOR\" commands needs upload file name")
				statusCode = StatusBadArguments
			}
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
		case command.TYPE:
			if len(args) == 1 || len(args) == 2 {
				statusCode, err = pi.dataType(args)
			} else {
				_, err = pi.printf("invalid argument length, TYPE takes 1 or 2 arguments. ")
				statusCode = StatusBadArguments
			}
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
	case command.PASV:
		conn, err = pi.listener.Accept()
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

	if !fi.info.IsDir() {
		pi.printf("%v: is not directory ", fi.path)
		return StatusBadArguments, nil
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
