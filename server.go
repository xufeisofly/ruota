package ruota

import (
	"bufio"
	"fmt"
)

type RServer struct {
	Socket    RServerSocket
	Processor RProcessor
}

func NewRServer(serverSocket *RServerSocket, processor RProcessor) (RServer, error) {
	return RServer{
		Socket:    *serverSocket,
		Processor: processor,
	}, nil
}

// 启动服务端 Socket 监听
func (p *RServer) Start() error {
	err := p.Listen()
	if err != nil {
		return err
	}
	p.AcceptLoop()
	return err
}

func (p *RServer) Listen() error {
	return p.Socket.Listen()
}

func (p *RServer) AcceptLoop() {
	fmt.Println("Server Socket Accept Loop")
	var b = make([]byte, 1024)
	for {
		p.Socket.Accept()

		n, err := bufio.NewReader(p.Socket.Conn).Read(b)
		data := string(b[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)

		// parse data
	}
}
