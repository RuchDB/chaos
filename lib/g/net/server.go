package net

import (
	"net"

	"github.com/RuchDB/chaos/lib/g/util"
	"github.com/RuchDB/chaos/lib/g/net/codec"
)

const (
	IPV4_ANY   = "0.0.0.0"
	IPV4_LOCAL = "127.0.0.1"

	SERVER_CONN_MAX = 1000000
)

/************************* TCP/IPv4 Server *************************/

const (
	SERVER_STATUS_UNINITED = 0
	SERVER_STATUS_INITED   = 1
	SERVER_STATUS_RUNNING  = 2
	SERVER_STATUS_STOPPED  = 3

	SERVER_CTRL_FLAG_INIT = 0
	SERVER_CTRL_FLAG_RUN  = 1
	SERVER_CTRL_FLAG_STOP = 2
)

type Server struct {
	addr     *net.TCPAddr
	listener *net.TCPListener
	status   int
	ctrlFlag int

	pktCodec codec.Codec
}

func CreateServer(addr string) (*Server, error) {
	taddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Server{
		addr: taddr,

		status:   SERVER_STATUS_UNINITED,
		ctrlFlag: SERVER_CTRL_FLAG_INIT,
	}, nil
}

func (server *Server) SetNetCodec(netCodec codec.Codec) {
	server.pktCodec = netCodec
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

/************************* Server Builder *************************/

type serverBuilder struct {
	listenIp   string
	listenPort uint16

	maxConns int
}

func NewServerBuilder() *serverBuilder {
	return &serverBuilder{
		listenIp: IPV4_ANY,
		listenPort: 2020,

		maxConns: 10000,
	}
}

func (builder *serverBuilder) SetListenAddr(ip string, port int) *serverBuilder {
	builder.listenIp = ip
	builder.listenPort = uint16(port)
	return builder
}

func (builder *serverBuilder) SetPacketCodec() *serverBuilder {
	return builder
}

func (builder *serverBuilder) SetMaxConns(conns int) *serverBuilder {
	conns = util.MaxInt(conns, 1)
	conns = util.MinInt(conns, SERVER_CONN_MAX)
	builder.maxConns = conns
	return builder
}

func (builder *serverBuilder) Build() *Server {
	return &Server{
		listenIp:   builder.listenIp,
		listenPort: builder.listenPort,
		
		connManager: NewConnectionManager(builder.maxConns),
	}
}
