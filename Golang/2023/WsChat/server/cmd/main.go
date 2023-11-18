package main

import (
	"WsChat/server/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/start", handlers.SocketHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
