package websocket

import (
	"time"

	jsonrpc2websocket "github.com/asuleymanov/jsonrpc2/websocket"
	"github.com/asuleymanov/websocket"
)

// ObjectStream implements jsonrpc2.ObjectStream that uses a WebSocket.
// It extends jsonrpc2/websocket.ObjectStream with read/write timeouts.
type ObjectStream struct {
	conn   *websocket.Conn
	stream jsonrpc2websocket.ObjectStream

	writeTimeout time.Duration
	readTimeout  time.Duration
}

//NewObjectStream initialised ObjectStream
func NewObjectStream(conn *websocket.Conn, writeTimeout, readTimeout time.Duration) *ObjectStream {
	return &ObjectStream{conn, jsonrpc2websocket.NewObjectStream(conn), writeTimeout, readTimeout}
}

//WriteObject data record in ObjectStream
func (stream *ObjectStream) WriteObject(v interface{}) error {
	stream.conn.SetWriteDeadline(time.Now().Add(stream.writeTimeout))
	return stream.stream.WriteObject(v)
}

//ReadObject reading data from ObjectStream
func (stream *ObjectStream) ReadObject(v interface{}) error {
	stream.conn.SetReadDeadline(time.Now().Add(stream.readTimeout))
	return stream.stream.ReadObject(v)
}

//Close closing the ObjectStream
func (stream *ObjectStream) Close() error {
	return stream.stream.Close()
}
