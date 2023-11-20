package main

import (
	"WsChat2/internal/service"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const url = "localhost:8080"

func main() {
	flag.Parse()
	hub := service.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		service.ServeWs(hub, w, r)
	})
	fmt.Printf("Server has been started at http://%s\n", url)
	log.Fatal(http.ListenAndServe(url, nil))
}
