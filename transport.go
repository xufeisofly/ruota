package ruota

import "io"

type RTransport interface {
	io.Reader
	io.Writer
	io.Closer
	Flush() error
}
