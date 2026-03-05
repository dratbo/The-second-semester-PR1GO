package http

import (
	"context"
	"net/http"
	"time"

	"tech-ip-sem2/shared/middleware"
	"tech-ip-sem2/services/tasks/internal/client/authclient"
	"tech-ip-sem2/services/tasks/internal/http/handlers"
)

type Server struct {
	server *http.Server
}

func NewServer(port, authBaseURL string) *Server {
	authClient := authclient.NewClient(authBaseURL, 3*time.Second)

	mux := http.NewServeMux()

	taskHandlers := handlers.NewTaskHandlers(authClient)

	mux.HandleFunc("POST /v1/tasks", taskHandlers.CreateTask)
	mux.HandleFunc("GET /v1/tasks", taskHandlers.ListTasks)
	mux.HandleFunc("GET /v1/tasks/{id}", taskHandlers.GetTask)
	mux.HandleFunc("PATCH /v1/tasks/{id}", taskHandlers.UpdateTask)
	mux.HandleFunc("DELETE /v1/tasks/{id}", taskHandlers.DeleteTask)
	
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
