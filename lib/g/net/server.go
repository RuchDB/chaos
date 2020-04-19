package net

import (
	"net"

	"github.com/RuchDB/chaos/log"
)

const (
	IPV4_ANY   = "0.0.0.0"
	IPV4_LOCAL = "127.0.0.1"

	SERVER_TYPE_TCP = "tcp"

	SERVER_CONN_MAX = 10000
)

type Server struct {
	listenIp   string
	listenPort uint16

	listener net.Listener

	connManager *ConnectionManager
}

func NewTcp4Server(ipv4 string, port uint16) *Server {
	return &Server{
		listenIp:   ipv4,
		listenPort: port,
	}
}

func (server *Server) Start() error {
	addr := IPv4AddrString(server.listenIp, server.listenPort)
	listener, err := net.Listen(SERVER_TYPE_TCP, addr)
	if err != nil {
		log.Errorf("Fail to start TCP server on [%s]. [%v]", addr, err)
		return err
	}
	log.Infof("Start TCP server on [%s]", addr)

	server.listener = listener
	server.connManager = NewConnectionManager(SERVER_CONN_MAX)

	// Loop to accept incoming connections
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Debugf("Fail to accept incoming connection. [%v]", err)
			continue
		}

		// Delegate the incoming connection to connection manager
		server.connManager.Handle(conn)
	}

	return nil
}
