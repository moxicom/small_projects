package main

import (
	"WsChat/client/internal"
	"log"

	"github.com/gorilla/websocket"
)

var done chan interface{}

func main() {
	done = make(chan interface{})

	socketUrl := "ws://localhost:8080" + "/start"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to websocket server:", err)
	}
	defer conn.Close()

	internal.ClearScreen()
	internal.RunChat(conn, done)
}
