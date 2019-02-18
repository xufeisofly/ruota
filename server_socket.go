package ruota

import (
	"net"
	"time"
)

type RServerSocket struct {
	Addr     net.Addr
	Conn     net.Conn
	Timeout  time.Duration
	Listener net.Listener
}

const DEFAULT_SERVER_TIMEOUT = 10 * time.Second

func NewRServerSocket(host, port string) (*RServerSocket, error) {
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
	return &RServerSocket{Addr: addr, Timeout: DEFAULT_SERVER_TIMEOUT}, err
}

func (p *RServerSocket) Listen() error {
	l, err := net.Listen("tcp", p.Addr.String())
	if err != nil {
		return err
	}
	// Save listener to RServerSocket
	p.Listener = l
	return nil
}

func (p *RServerSocket) Accept() error {
	conn, err := p.Listener.Accept()
	if err != nil {
		return err
	}
	p.Conn = conn
	return nil
}
