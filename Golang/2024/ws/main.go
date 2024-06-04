package main

import (
	"fmt"
	"log"
	"net/http"
	"ws/internal/ws"
)

func main() {

	hub := ws.NewHub()

	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.Upgrade(hub, w, r)
	})

	fmt.Println("Server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
