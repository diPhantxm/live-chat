package app

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type Server struct {
	conn     net.Conn
	listener net.Listener
	clients  map[net.Addr]net.Conn
	config   *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		conn:     nil,
		listener: nil,
		clients:  map[net.Addr]net.Conn{},
		config:   config,
	}
}

func (s *Server) Start() {
	var err error
	s.listener, err = net.Listen(s.config.Protocol, fmt.Sprintf("%s:%d", s.config.Address, s.config.Port))

	if err != nil {
		log.Fatalf("Error occured when starting listening: %s", err.Error())
	}

	defer s.listener.Close()

	for {
		var incConn net.Conn
		incConn, err = s.listener.Accept()
		if err != nil {
			log.Fatalf("Error occurred when accepting message: %s", err.Error())
		}

		s.clients[incConn.RemoteAddr()] = incConn
		s.SendToOther(incConn, fmt.Sprintf("New client connected. Address: %s\r\n", incConn.RemoteAddr().String()))
		go s.broadcast(incConn)
	}
}

func (s *Server) broadcast(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')

		if err == io.EOF {
			delete(s.clients, conn.RemoteAddr())
			s.SendToOther(conn, fmt.Sprintf("Client disconnected. Address: %s\r\n", conn.RemoteAddr().String()))
			return
		}

		if err != nil {
			log.Fatalf("Could not read message: %s\r\n", err.Error())
		}

		s.SendToOther(conn, fmt.Sprintf("[%s]: %s", conn.RemoteAddr(), msg))
	}
}

func (s *Server) SendToOther(current net.Conn, msg string) {
	for _, client := range s.clients {
		if client.RemoteAddr() != current.RemoteAddr() {
			client.Write([]byte(msg))
		}
	}
}

func (s *Server) Close() {
	if err := s.conn.Close(); err != nil {
		log.Fatalf("Error occurred when closing connection: %s\n", err.Error())
	}
}
