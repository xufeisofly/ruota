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

func (p *RClientSocket) Read(b []byte) (int, error) {
	return p.Conn.Read(b)
}

func (p *RClientSocket) Write(b []byte) (int, error) {
	return p.Conn.Write(b)
}

func (p *RClientSocket) Flush() error {
	return nil
}
