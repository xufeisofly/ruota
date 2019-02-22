package ruota

import "io"

type RTransport interface {
	io.Reader
	io.Writer
	io.Closer
	Flush() error

	WriteFunName([]byte) error
	WriteArg([]byte) error
	WriteByte(byte) error
	WriteList([][]byte, RType) error

	ReadFunName() ([]byte, error)
	ReadByte() (byte, error)
	ReadList() ([][]byte, int, error)
}
