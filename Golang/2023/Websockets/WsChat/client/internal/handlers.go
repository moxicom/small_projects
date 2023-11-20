package internal

import (
	"log"

	"github.com/gorilla/websocket"
)

func receiveHandler(connection *websocket.Conn, done chan interface{}) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		PrintColoredText(string(msg), "\033[34m")
	}
}
