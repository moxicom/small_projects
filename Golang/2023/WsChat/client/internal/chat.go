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

func RunChat(conn *websocket.Conn, done chan interface{}) {

	interrupt = make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

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
