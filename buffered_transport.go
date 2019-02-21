package ruota

import "bufio"

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
	return p.ReadWriter.Flush()
}

// Write and Read Methods

func (p *RBufferedTransport) WriteFunName(funName []byte) error {
	_, err := p.Write(funName)
	return err
}

func (p *RBufferedTransport) WriteArg(arg []byte) error {
	_, err := p.Write(arg)
	return err
}

func (p *RBufferedTransport) WriteList(l [][]byte, elemType RType) error {
	// 写入数组元素类型
	if _, err := p.Write(elemType); err != nil {
		return err
	}
	// 写入数组大小
	if _, err := p.Write(len(l)); err != nil {
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
	_, err := p.Read(b)
	return b, err
}

func (p *RBufferedTransport) ReadList() ([][]byte, size, error) {
	// 读取数组元素类型
	var rType RType
	if _, err := p.Read(rType); err != nil {
		return [][]byte{}, 0, err
	}
	// 读取数组大小
	var size int32
	if _, err := p.Read(size); err != nil {
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
