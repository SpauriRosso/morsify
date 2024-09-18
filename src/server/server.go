package server

import (
	"net/http"
	"time"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Start(addr string) error {
	serv := &http.Server{
		Addr:              addr,
		Handler:           s.mux,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1MB
	}
	key := "src/certfiles/dev/localhost-key.pem"
	certFile := "src/certfiles/dev/localhost.pem"
	return serv.ListenAndServeTLS(certFile, key)
}
