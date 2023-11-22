package main

import (
	"WsChat2/internal/service"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Print("Введите ваше имя: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Ошибка чтения имени:", err)
	}
	username = strings.TrimSpace(username) // Убираем пробелы и символы новой строки

	socketUrl := "ws://localhost:8080" + "/ws?username=" + username
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
