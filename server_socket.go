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

func (p *RServerSocket) Close() error {
	return p.Conn.Close()
}

func (p *RServerSocket) Flush() error {
	return nil
}

func (p *RServerSocket) Write(b []byte) (int, error) {
	return p.Conn.Write(b)
}

func (p *RServerSocket) Read(b []byte) (int, error) {
	return p.Conn.Read(b)
}

func (p *RServerSocket) ReadByte() (byte, error) {
	return [1]byte{0}[0], nil
}

func (p *RServerSocket) ReadFunName() ([]byte, error) {
	return []byte{}, nil
}

func (p *RServerSocket) ReadList() ([][]byte, int, error) {
	return [][]byte{}, 0, nil
}

func (p *RServerSocket) WriteFunName([]byte) error {
	return nil
}

func (p *RServerSocket) WriteArg([]byte) error {
	return nil
}

func (p *RServerSocket) WriteByte(byte) error {
	return nil
}

func (p *RServerSocket) WriteList([][]byte, RType) error {
	return nil
}
