package main

import (
	"WsChat/client/internal"
	"bufio"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var done chan interface{}
var interrupt chan os.Signal

func receivingHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		internal.PrintColoredText(string(msg), "\033[34m")
	}
}

func userInput(inputChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	inputChan <- msg
}

func main() {
	done = make(chan interface{})
	interrupt = make(chan os.Signal)

	signal.Notify(interrupt, os.Interrupt)

	socketUrl := "ws://localhost:8080" + "/handshake"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to websocket server:", err)
	}
	internal.ClearScreen()
	defer conn.Close()

	go receivingHandler(conn)

	for {
		inputChan := make(chan string)
		go userInput(inputChan)
		select {
		case input := <-inputChan:
			err = conn.WriteMessage(websocket.TextMessage, []byte(input))
			if err != nil {
				log.Println("Error during writing to websocket:", err)
				return
			}
			internal.PrintColoredText("message has been sent", "\033[32m")

		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		}
	}
}
