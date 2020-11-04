package server

import (
	"log"
	"net/http"
)

type Server struct {
	location string
	port string
}

func New(location string, port string) *Server {
	return &Server{location:location,  port: port}
}

func (s *Server) Start() error {
	fs := http.FileServer(http.Dir("./"+s.location))
	http.Handle("/", fs)

	log.Println("Listening on http://localhost"+s.port)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}

	return nil
}