package main

import (
	"github.com/linusbohwalli/feature-toggle-service/http-api"
	"log"
	"net"
	"net/http"
)

const DefaultAddr = ":9091"

type Server struct {
	ln      net.Listener
	handler *http_api.Handler

	Addr string
}

func NewServer() *Server {
	return &Server{
		Addr: DefaultAddr,
	}
}

func (s *Server) Open() error {

	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	s.ln = ln

	if err := http.Serve(s.ln, s.handler); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func main() {

	s := NewServer()

	h := http_api.NewHandler()
	s.handler = h
	if err := s.Open(); err != nil {
		log.Fatal(err)
	}
}
