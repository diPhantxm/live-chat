package app

import (
	"bufio"
	"log"
	"net"
)

type Server struct {
	conn     net.Conn
	listener net.Listener
	clients  []net.Addr
}

func NewServer() *Server {
	return &Server{
		conn:     nil,
		listener: nil,
	}
}

func (s *Server) Start() {
	var err error
	s.listener, err = net.Listen("tcp", ":8080")

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

		s.clients = append(s.clients, s.conn.RemoteAddr())
		go s.broadcast(incConn)
	}
}

func (s *Server) broadcast(incConn net.Conn) {
	msg, err := bufio.NewReader(s.conn).ReadString('\n')
	if err != nil {
		log.Fatalf("Could not read message: %s", err.Error())
	}

	for _, client := range s.clients {
		if client != incConn.RemoteAddr() {
			incConn.Write([]byte(msg))
		}
	}
}

func (s *Server) Close() {
	if err := s.conn.Close(); err != nil {
		log.Fatalf("Error occurred when closing connection: %s\n", err.Error())
	}
}
