package ruota

import "io"

type RTransport interface {
	io.Reader
	io.Writer
	io.Closer
	Flush() error

	WriteFunName([]byte) error
	WriteArg([]byte) error
	WriteList([][]byte, RType) error

	ReadFunName() ([]byte, error)
	ReadList() ([][]byte, size, error)
}
