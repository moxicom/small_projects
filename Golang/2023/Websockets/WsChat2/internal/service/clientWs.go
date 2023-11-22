package service

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type ClientCon struct {
	con *websocket.Conn
	// Channel of outbound messages
	send chan []byte
}

func NewClient(con *websocket.Conn, send chan []byte) *ClientCon {
	return &ClientCon{con: con, send: send}
}

func (c *ClientCon) ReceiveMessage() {
	con := c.con
	defer con.Close()
	con.SetPingHandler(
		func(appData string) error {
			// fmt.Println()
			// log.Printf("PING received\n")
			con.SetWriteDeadline(time.Now().Add(writeWait))
			err := con.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(writeWait))
			if err != nil {
				return err
			}
			return nil
		})
	for {
		_, data, err := con.ReadMessage()
		if err != nil {
			return
		}
		log.Println(string(data))
	}
}

func (c *ClientCon) WriteMessage() {
	defer c.con.Close()

	for {
		go GetMessageFromConsole(c.send)
		select {
		case message, ok := <-c.send:
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
		}
	}
}
