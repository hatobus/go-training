package interpreter

import (
	"fmt"
	"net"
	"time"
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
	pi.conn.Write([]byte(fmt.Sprintf("HI! %v\n", time.Now().Format(time.RFC822))))
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
