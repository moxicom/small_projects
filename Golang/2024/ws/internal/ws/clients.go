package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 5 * time.Second
	pongWait       = 5 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	ID   uint64
	Hub  *Hub
	Con  *websocket.Conn
	Send chan Message
}

func (c *Client) readWS() {
	defer func() {
		c.Hub.Unregister <- c
		c.Con.Close()
	}()

	c.Con.SetReadLimit(maxMessageSize)
	c.Con.SetReadDeadline(time.Now().Add(pongWait))
	c.Con.SetPongHandler(func(appData string) error {
		c.Con.SetReadDeadline(time.Now().Add(pongWait))
		log.Println("Pong handler")
		return nil
	})

	for {
		_, byteMsg, err := c.Con.ReadMessage()

		log.Println("Accepted a message")

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err.Error())
			}
			break
		}

		var msg Message

		err = json.Unmarshal(byteMsg, &msg)
		if err != nil {
			log.Printf("error %v\n", err.Error())
			break
		}

		c.Hub.Broadcast <- msg
	}
}

func (c *Client) writeWS() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		c.Hub.Unregister <- c
		c.Con.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.Con.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				log.Println("Send chan is closed")
				c.Con.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Con.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err.Error())
				return
			}

			w.Write([]byte(msg.msg))

			if err := w.Close(); err != nil {
				log.Println(err.Error())
				return
			}
		case <-ticker.C:
			c.Con.SetWriteDeadline(time.Now().Add(writeWait))
			log.Println("Sending ping msg")
			if err := c.Con.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}
