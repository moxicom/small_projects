package main

import (
	"WsChat/server/internal/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/start", service.SocketHandler)
	http.HandleFunc("/", service.GoHome)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
