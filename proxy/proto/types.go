package proto

import "errors"

// ReqType request type
type ReqType int8

// Request type
const (
	ReqRead ReqType = iota
	ReqWrite
)

// Route handlers
const (
	AllAsynsRoute = "AllAsync"
	AllSync       = "AllSync"
)

// Route type
const ()

// response error
var (
	ErrKeyNotFound = errors.New("key not found")
)

// Request request interface.
type Request interface {
	CmdString() string
	Cmd() []byte
	Key() []byte
	Put()
	ReqType() ReqType
}

// ProxyConn decode bytes from client and encode write to conn.
type ProxyConn interface {
	Decode([]*Message) ([]*Message, error)
	Encode(msg *Message) error
	Flush() error
}

// NodeConn handle Msg to backend cache server and read response.
type NodeConn interface {
	Write(*Message) error
	Read(*Message) error
	Flush() error
	Close() error
}

// Pinger for executor ping node.
type Pinger interface {
	Ping() error
	Close() error
}

// Forwarder is the interface for backend run and process the messages.
type Forwarder interface {
	Forward([]*Message) error
	Close() error
}
