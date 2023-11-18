package internal

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var interrupt chan os.Signal

func userInput(inputChan chan string) {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	inputChan <- msg
}

func chatInputHandler(conn *websocket.Conn, done chan interface{}) {
	for {
		inputChan := make(chan string)
		go userInput(inputChan)
		select {
		case input := <-inputChan:
			if strings.TrimSpace(input) == "" {
				continue
			}
			err := conn.WriteMessage(websocket.TextMessage, []byte(input))
			if err != nil {
				log.Println("Error during writing to websocket:", err)
				return
			}
			PrintColoredText("message has been sent", "\033[32m")

		case <-interrupt:
			log.Println("Interrupt signal received. Closing all pending connections")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel")
			}
			return
		}
	}
}

func RunChat(conn *websocket.Conn, done chan interface{}) {
	go receiveHandler(conn, done)

	interrupt = make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	chatInputHandler(conn, done)
}
