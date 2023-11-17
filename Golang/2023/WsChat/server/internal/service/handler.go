package service

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options
var connections = []*websocket.Conn{}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()
	connections = append(connections, conn)
	log.Println(connections)
	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			connections = removeConn(connections, conn)
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Server: %s", message)
		for _, c := range connections {
			if c != conn {
				err = c.WriteMessage(messageType, message[:len(message)-1])
				if err != nil {
					connections = removeConn(connections, conn)
					log.Println("Error during message writing:", err)
					break
				}
			}
		}
	}
	connections = removeConn(connections, conn)
}

func removeConn(slice []*websocket.Conn, val *websocket.Conn) []*websocket.Conn {
	index := -1
	for i, v := range slice {
		if v == val {
			index = i
			break
		}
	}

	// If the element is found, remove it using slicing
	if index != -1 {
		// Ensure that the index is within the bounds of the slice
		if index < len(slice)-1 {
			// Move the elements after the removed one to the left
			copy(slice[index:], slice[index+1:])
		}

		// Resize the slice to remove the last element
		slice = slice[:len(slice)-1]
	}

	return slice
}

func GoHome(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Index page")
}
