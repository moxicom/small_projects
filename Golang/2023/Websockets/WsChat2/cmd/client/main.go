package main

import (
	"WsChat2/internal/service"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	socketUrl := "ws://localhost:8080" + "/ws"
	con, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to websocket server", err)
	}
	defer con.Close()

	inputChan := make(chan []byte)
	client := service.NewClient(con, inputChan)
	go client.ReceiveMessage()
	client.WriteMessage()
}
