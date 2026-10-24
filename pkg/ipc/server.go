package ipc

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Server struct {
	socketPath string
	listener   net.Listener
}

func NewServer(socketPath string) *Server {
	return &Server{socketPath: socketPath}
}

func (s *Server) Start(handler func(req string) string) error {
	os.Remove(s.socketPath)

	l, err := net.Listen("unix", s.socketPath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écoute IPC: %w", err)
	}
	s.listener = l

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn, handler)
	}
}

func (s *Server) handleConnection(conn net.Conn, handler func(req string) string) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		req := scanner.Text()
		resp := handler(req)
		conn.Write([]byte(resp + "\n"))
	}
}

func (s *Server) Stop() {
	if s.listener != nil {
		s.listener.Close()
	}
	os.Remove(s.socketPath)
}
