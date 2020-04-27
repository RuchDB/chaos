package net

import (
	"net"

	"github.com/RuchDB/chaos/buf"
	"github.com/RuchDB/chaos/log"
	"github.com/RuchDB/chaos/util"
	"github.com/RuchDB/chaos/types"
)

/************************* Connection *************************/

const (
	CONN_RBUF_MIN_SIZE = 128
	CONN_RBUF_MAX_SIZE = 1024 * 1024 * 16
)

type Connection struct {
	conn net.Conn
	buff buf.Buffer

	tsCreated    int64
	tsLastActive int64
}

func NewConnection(conn net.Conn) *Connection {
	tsCur := util.Timestamp(util.Now())

	return &Connection{
		conn: conn,
		buff: buf.NewDecodeBuffer(CONN_RBUF_MIN_SIZE, CONN_RBUF_MAX_SIZE),

		tsCreated:    tsCur,
		tsLastActive: tsCur,
	}
}

func (conn *Connection) OpenSession() {
	
}

/************************* Connection Manager *************************/

type ConnectionManager struct {
	conns    *types.Set // a set of connections
	maxConns int
}

func NewConnectionManager(maxConns int) *ConnectionManager {
	return &ConnectionManager{
		conns:    types.NewSetWithInitSize(maxConns),
		maxConns: maxConns,
	}
}

func (manager *ConnectionManager) Handle(tcpConn net.Conn) {
	// Reach max connection size, reject new incomming connections
	if manager.conns.Len() >= manager.maxConns {
		log.Warnf("Reach max connection size [%d], reject incoming one [%s]", 
			manager.maxConns, tcpConn.RemoteAddr().String())
		return
	}

	conn := NewConnection(tcpConn)
	manager.conns.Insert(conn)

	go conn.OpenSession()
}
