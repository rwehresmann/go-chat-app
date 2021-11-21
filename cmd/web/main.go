package main

import (
	"log"
	"net/http"

	"github.com/rwehresmann/chat-app/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener...")
	go handlers.ListenToWsChannel()

	log.Println("Starting server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
