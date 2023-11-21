package service

import (
	"bufio"
	"os"
)

func GetMessageFromConsole(inputChan chan []byte) {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	inputChan <- []byte(msg[:len(msg)-1])
}
