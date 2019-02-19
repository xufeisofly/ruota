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
