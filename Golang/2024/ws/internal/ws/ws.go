package ws

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

// TODO accept message

// TODO send message

var upgrader = websocket.Upgrader{
	// Reuse buffers that allocates standart library http server
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

func Upgrade(hub *Hub, w http.ResponseWriter, r *http.Request) {
	log.Println("Handled WS")

	userIDstr := r.URL.Query().Get("id")
	if userIDstr == "" {
		http.Error(w, fmt.Errorf("No id in query").Error(), http.StatusBadRequest)
		fmt.Printf("error: %s\n", "No id in query")
		return
	}

	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("User ID on upgrade %v\n", userID)

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	client := &Client{ID: userID, Hub: hub, Con: con, Send: make(chan Message, 256)}
	client.Hub.Register <- client

	go client.readWS()
	go client.writeWS()
}
