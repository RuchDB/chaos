package net

import (
	"fmt"
	"net"
	"time"
	"errors"

	"github.com/RuchDB/chaos/lib/g/util"
	"github.com/RuchDB/chaos/lib/g/net/codec"
)

const (
	IPV4_ANY   = "0.0.0.0"
	IPV4_LOCAL = "127.0.0.1"

	SERVER_CONN_MAX = 1000000

	SERVER_ACCEPT_TIMEOUT_MS = 100
)

/************************* TCP/IPv4 Server *************************/

const (
	SERVER_STATUS_INITED   = 0
	SERVER_STATUS_RUNNING  = 1
	SERVER_STATUS_STOPPED  = 2

	SERVER_CTRL_FLAG_INIT = 0
	SERVER_CTRL_FLAG_RUN  = 1
	SERVER_CTRL_FLAG_STOP = 2
)

type Server struct {
	addr     *net.TCPAddr
	listener *net.TCPListener
	status   int
	ctrlFlag int

	netCodec codec.Codec

	connManager *ConnectionManager
}

func NewServer(addr *net.TCPAddr, netCodec codec.Codec, maxConns int) *Server {
	return &Server{
		addr:     addr,
		status:   SERVER_STATUS_INITED,
		ctrlFlag: SERVER_CTRL_FLAG_INIT,
		
		netCodec: netCodec,

		connManager: NewConnectionManager(maxConns),
	}
}

func (server *Server) Run() error {
	logger.Infof("Start to listen on [%s]", server.addr.String())
	listener, err := net.ListenTCP("tcp", server.addr)
	if err != nil {
		logger.Errorf("Listen on [%s] error: [%v]", server.addr.String(), err)
		return fmt.Errorf("Fail to listen on TCP [%s]: [%v]", server.addr.String(), err)
	}

	server.listener = listener
	server.ctrlFlag = SERVER_CTRL_FLAG_RUN
	server.status = SERVER_STATUS_RUNNING

	// Accept timeout
	server.listener.SetDeadline(time.Now().Add(time.Millisecond * SERVER_ACCEPT_TIMEOUT_MS))

	logger.Infof("Listen on TCP [%s] & Wait for incoming connections", server.addr.String())
	// Loop to accept incoming connections
	var serr error
	for server.status == SERVER_STATUS_RUNNING && server.ctrlFlag != SERVER_CTRL_FLAG_STOP {
		conn, err := server.listener.AcceptTCP()
		if err != nil {
			if IsTimeoutError(err) || IsTemporaryError(err) {
				continue
			}

			logger.Errorf("Accept connections error: [%v]", err)
			serr = fmt.Errorf("Internal server error: [%v]", err)
			break
		}
		logger.Infof("Remote connection [%s] is accepted", conn.RemoteAddr().String())

		// Delegate the incoming connection to connection manager
		server.connManager.Handle(conn)
	}
	
	logger.Infof("Shutdown server on TCP [%s]", server.addr.String())
	server.status = SERVER_STATUS_STOPPED
	server.listener.Close()
	// Close/Clear remote connections
	server.connManager.Clear()

	return serr
}

func (server *Server) ListenAddr() net.Addr {
	return server.addr
}

func (server *Server) IsRunning() bool {
	return server.status == SERVER_STATUS_RUNNING
}

func (server *Server) Stop() {
	server.ctrlFlag = SERVER_CTRL_FLAG_STOP
}


/************************* Server Creator *************************/

type ServerCreator struct {
	addr     *net.TCPAddr
	netCodec codec.Codec

	maxConns int
}

func NewServerCreator() *ServerCreator {
	return &ServerCreator{
		maxConns: 10000,
	}
}

func (creator *ServerCreator) ConfigListenAddr(addr *net.TCPAddr) error {
	if addr == nil {
		return errors.New("config error: nil address")
	}

	creator.addr = addr
	return nil
}

func (creator *ServerCreator) ConfigListenAddrByString(addr string) error {
	addr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return errors.New("config error: invalid address [ipv4:port]")
	}

	creator.addr = addr
	return nil
}

func (creator *ServerCreator) ConfigListenAddrByIpPort(ip string, port int) error {
	return creator.ConfigListenAddrByString(fmt.Sprintf("%s:%d", ip, port))
}

func (creator *ServerCreator) ConfigMaxConns(conns int) error {
	conns = util.MaxInt(conns, 1)
	conns = util.MinInt(conns, SERVER_CONN_MAX)
	creator.maxConns = conns
	return nil
}

func (creator *ServerCreator) ConfigNetCodec(c codec.Codec) error {
	if c == nil {
		return errors.New("config error: nil net codec")
	}

	creator.netCodec = c
	return nil
}

func (creator *ServerCreator) ConfigNetCodecById(cid int) error {
	c, err := codec.GetCodecById(cid)
	if err != nil {
		return errors.New("config error: invalid net codec id")
	}

	creator.netCodec = c
	return nil
}

func (creator *ServerCreator) ConfigNetCodecByName(cname string) error {
	c, err := codec.GetCodecByName(cname)
	if err != nil {
		return errors.New("config error: invalid net codec name")
	}

	creator.netCodec = c
	return nil
}

func (creator *ServerCreator) CreateServer() (*Server, error) {
	if creator.addr == nil {
		return nil, errors.New("create server failed: nil address")
	}
	if creator.netCodec == nil {
		return nil, errors.New("create server failed: nil net codec")
	}

	return NewServer(creator.addr, creator.netCodec, creator.maxConns), nil
}
