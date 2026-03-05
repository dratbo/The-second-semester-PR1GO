package http

import (
	"context"
	"net/http"
	"time"

	"tech-ip-sem2/services/auth/internal/http/handlers"
	"tech-ip-sem2/shared/middleware"
)

type Server struct {
	server *http.Server
}

func NewServer(port string) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/auth/login", handlers.Login)
	mux.HandleFunc("GET /v1/auth/verify", handlers.Verify)

	handler := middleware.RequestID(mux)
	handler = middleware.Logging(handler)

	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
