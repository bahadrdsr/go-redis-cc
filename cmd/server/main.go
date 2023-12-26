package main

import (
	"log"
	"net"

	"github.com/bahadrdsr/go-redis-cc/internal/app/handler"
	"github.com/bahadrdsr/go-redis-cc/internal/store"
)

func main() {
	dataStore := store.New()
	commandHandler := handler.New(dataStore)
	port := "6379"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	log.Printf("Server started on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handler.HandleConnection(conn, commandHandler)
	}
}
