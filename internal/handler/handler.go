package handler

import "net/http"

type Server struct {
	server *http.Server
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func NewServer(port, host string) *Server {
	if port == "" {
		port = "8080"
	}

	return &Server{
		server: &http.Server{
			Addr:    host + ":" + port,
			Handler: Route(),
		},
	}
}
