package ruota

import (
	"fmt"
)

type RServer struct {
	Trans      RTransport
	Processor  RProcessor
	Serializer RSerializer
}

func NewRServer(trans RTransport, processor RProcessor, serializer RSerializer) (RServer, error) {
	return RServer{
		Trans:      trans,
		Processor:  processor,
		Serializer: serializer,
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
	return p.Trans.Listen()
}

func (p *RServer) AcceptLoop() {
	fmt.Println("Server Socket Accept Loop")
	for {
		p.Trans.Accept()

		err := p.Processor.Call(p.Trans, p.Serializer)
		if err != nil {
			return
		}
	}
}
