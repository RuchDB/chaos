package net

import (
	"net"

	"github.com/RuchDB/chaos/lib/g/buf"
	"github.com/RuchDB/chaos/lib/g/util"
	"github.com/RuchDB/chaos/lib/g/net/codec"
)


/************************* Connection *************************/

const (
	CONN_RBUF_MIN_SIZE = 128
	CONN_RBUF_MAX_SIZE = 1024 * 1024 * 16
)

type Connection struct {
	conn *net.TCPConn

	buff buf.Buffer

	netCodec codec.Codec

	tsCreated    int64
	tsLastActive int64
}

func NewConnection(conn *net.TCPConn, netCodec codec.Codec) *Connection {
	return &Connection{
		conn: conn,

		buff: buf.NewDecodeBuffer(CONN_RBUF_MIN_SIZE, CONN_RBUF_MAX_SIZE),

		netCodec: netCodec,

		tsCreated:    util.TimestampMs(),
		tsLastActive: util.TimestampMs(),
	}
}


/************************* Connection Manager *************************/

type ConnectionManager struct {
}



