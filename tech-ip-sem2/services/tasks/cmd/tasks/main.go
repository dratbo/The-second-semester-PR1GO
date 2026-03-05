package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"tech-ip-sem2/services/tasks/internal/http"
)

func main() {
	port := os.Getenv("TASKS_PORT")
	if port == "" {
		port = "8082"
	}

	authBaseURL := os.Getenv("AUTH_BASE_URL")
	if authBaseURL == "" {
		authBaseURL = "http://localhost:8081"
	}

	srv := http.NewServer(port, authBaseURL)

	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("Server stopped: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down tasks service...")
	if err := srv.Stop(); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
}
