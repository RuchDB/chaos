package net

import (
	"net"

	"github.com/RuchDB/chaos/lib/g/buf"
	"github.com/RuchDB/chaos/lib/g/util"
	"github.com/RuchDB/chaos/lib/g/types"
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

func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) Close() {

}


/************************* Connection Manager *************************/

type ConnectionManager struct {
	// Connection Map -- [remote addr --> conn]
	conns *types.ConcurrentMap
	
	maxConns int
}

func NewConnectionManager(maxConns int) *ConnectionManager {
	return &ConnectionManager{
		conns: types.NewConcurrentMap(),
		maxConns: maxConns,
	}
}

func (manager *ConnectionManager) Handle(c *Connection) {
	// Reach max connection size, reject new incomming connections
	if manager.conns.Len() >= manager.maxConns {
		logger.Warnf("Reach connection limit [%d], reject connection [%s]", 
			manager.maxConns, c.RemoteAddr().String())
		return
	}

	// If connection already exists, close previous instance first
	if val, exist := manager.conns.Get(c.RemoteAddr().String()); exist {
		if conn, ok := val.(*Connection); ok {
			conn.Close()
		}
		manager.conns.Delete(c.RemoteAddr().String())
	}

	// TODO: Open session for connection

	manager.conns.Put(c.RemoteAddr().String(), c)
}
