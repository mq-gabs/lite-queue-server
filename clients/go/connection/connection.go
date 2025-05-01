package connection

import (
	"bufio"
	"fmt"
	"net"
)

const (
	Host = "127.0.0.1"
	Port = 6633
)

type Connection struct {
	conn   net.Conn
	reader bufio.Reader
}

func New() (*Connection, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%v", Host, Port))

	if err != nil {
		return nil, fmt.Errorf("cannot create connection: %v", err)
	}

	return &Connection{
		conn:   conn,
		reader: *bufio.NewReader(conn),
	}, nil
}

func (c *Connection) Close() {
	c.conn.Close()
}

func (c *Connection) Handle() {

}
