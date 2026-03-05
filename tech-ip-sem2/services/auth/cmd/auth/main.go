package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"tech-ip-sem2/services/auth/internal/http"
)

func main() {
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8081"
	}

	srv := http.NewServer(port)

	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("Error starting HTTP server(дудосят): %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	if err := srv.Stop(); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}

}
