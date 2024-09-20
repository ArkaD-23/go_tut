package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("loop error:", err)
			continue
		}

		fmt.Println("New connection to the server:", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 2048)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		msg := buff[:n]
		fmt.Println("Message:", string(msg))
	}
}

func main() {
	server := NewServer(":3000")
	log.Fatal(server.Start())
}
