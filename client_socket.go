package ruota

import (
	"fmt"
	"net"
	"time"
)

type RClientSocket struct {
	Addr    net.Addr
	Conn    net.Conn
	Timeout time.Duration
}

const DEFAULT_TIMEOUT = 10 * time.Second

func NewRClientSocket(host, port string) (*RClientSocket, error) {
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
	return &RClientSocket{Addr: addr, Timeout: DEFAULT_TIMEOUT}, err
}

func (p *RClientSocket) Dial() error {
	conn, err := net.Dial("tcp", p.Addr.String())
	if err != nil {
		return err
	}
	p.Conn = conn
	fmt.Println("Client Socket Dialed")
	return err
}

func (p *RClientSocket) Close() error {
	return p.Conn.Close()
}
