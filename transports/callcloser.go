package transports

import "io"

//CallCloser network transport interface
type CallCloser interface {
	Caller
	io.Closer
}
