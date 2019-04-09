package ruota

import (
	"bufio"
	"fmt"
	"io"
)

type RBufferedTransport struct {
	ReadWriter bufio.ReadWriter
	Trans      RTransport
}

func NewRBufferedTransport(trans RTransport, bufferSize int) (RTransport, error) {
	return &RBufferedTransport{
		ReadWriter: bufio.ReadWriter{
			Reader: bufio.NewReaderSize(trans, bufferSize),
			Writer: bufio.NewWriterSize(trans, bufferSize),
		},
		Trans: trans,
	}, nil
}

// 在socket 的 RWC 基础上，重新定义了 BufferedTransport的 RWC 方法

func (p *RBufferedTransport) Close() error {
	return p.Trans.Close()
}

func (p *RBufferedTransport) Read(b []byte) (int, error) {
	n, err := p.ReadWriter.Read(b)
	if err != nil {
		p.ReadWriter.Reader.Reset(p.Trans)
	}
	return n, nil
}

func (p *RBufferedTransport) Write(w []byte) (int, error) {
	n, err := p.ReadWriter.Write(w)
	if err != nil {
		p.ReadWriter.Writer.Reset(p.Trans)
	}
	return n, nil
}

func (p *RBufferedTransport) Flush() error {
	if err := p.ReadWriter.Flush(); err != nil {
		p.ReadWriter.Writer.Reset(p.Trans)
		return err
	}
	return p.Trans.Flush()
}

// Write and Read Methods

func (p *RBufferedTransport) writeEnd() error {
	_, err := p.Write([]byte("\n"))
	return err
}

func (p *RBufferedTransport) WriteFunName(funName []byte) error {
	fmt.Printf("WriteFunName --> %s \n", string(funName))
	_, err := p.Write(funName)
	if err != nil {
		return err
	}
	err = p.writeEnd()
	return err
}

func (p *RBufferedTransport) WriteArg(arg []byte) error {
	_, err := p.Write(arg)
	if err != nil {
		return err
	}
	err = p.writeEnd()
	return err
}

func (p *RBufferedTransport) WriteByte(b byte) error {
	v := [1]byte{b}
	_, err := p.Write(v[0:1])
	return err
}

func (p *RBufferedTransport) WriteList(l [][]byte, elemType RType) error {
	// 写入数组元素类型
	if err := p.WriteByte(byte(elemType)); err != nil {
		return err
	}
	// 写入数组大小
	if err := p.WriteByte(byte(len(l))); err != nil {
		return err
	}
	// 写入数组内容
	for _, v := range l {
		if _, err := p.Write(v); err != nil {
			return err
		}
	}
	return nil
}

func (p *RBufferedTransport) ReadFunName() ([]byte, error) {
	var b []byte
	_, err := p.Read(b)
	return b, err
}

func (p *RBufferedTransport) ReadByte() (byte, error) {
	v := [1]byte{0}
	n, err := p.Read(v[0:1])
	if n > 0 && (err == nil || err == io.EOF) {
		return v[0], nil
	}
	if n > 0 && err != nil {
		return v[0], err
	}
	if err != nil {
		return 0, err
	}
	return v[0], nil
}

func (p *RBufferedTransport) ReadList() ([][]byte, int, error) {
	// 读取数组元素类型
	_, err := p.ReadByte()
	if err != nil {
		return [][]byte{}, 0, err
	}
	// 读取数组大小
	sizeByte, err := p.ReadByte()
	size := int(sizeByte)
	if err != nil {
		return [][]byte{}, 0, err
	}
	// 读取数据内容
	var ret [][]byte
	for i := 0; i < size; i++ {
		var _elem []byte
		if _, err := p.Read(_elem); err != nil {
			return [][]byte{}, 0, err
		}
		ret = append(ret, _elem)
	}
	return ret, size, nil
}

func (p *RBufferedTransport) Listen() error {
	return p.Trans.Listen()
}

func (p *RBufferedTransport) Accept() error {
	return p.Trans.Accept()
}
