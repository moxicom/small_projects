package service

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 5 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub *Hub

	// The websocket connection
	con *websocket.Conn

	// Buffered channel of outbound messages
	send chan []byte

	username string
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.con.Close()
	}()
	c.con.SetReadLimit(maxMessageSize)
	c.con.SetReadDeadline(time.Now().Add(pongWait))
	c.con.SetPongHandler(func(string) error {
		c.con.SetReadDeadline(time.Now().Add(pongWait))
		log.Println(" <- PONG")
		return nil
	})
	for {
		_, message, err := c.con.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
		message = []byte(c.username + ": " + string(message))
		c.hub.broadcast <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.con.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			// Set write limitation to determine the maximum time during which a server or client attempts to write data to a WebSocket channel
			c.con.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.con.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.con.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			log.Println("PING ->")
			c.con.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.con.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(h *Hub, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: h, con: con, send: make(chan []byte, 256), username: username}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()

}
