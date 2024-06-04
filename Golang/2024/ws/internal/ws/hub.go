package ws

import "sync"

type Hub struct {
	Clients    sync.Map
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

type Message struct {
	ToID uint64
	msg  string
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients.Store(client.ID, client)
		case client := <-h.Unregister:
			h.Clients.Delete(client.ID)
			close(client.Send)
		case msg := <-h.Broadcast:
			v, ok := h.Clients.Load(msg.ToID)
			if !ok {
				continue
			}

			targetClient := v.(Client)

			select {
			case targetClient.Send <- msg:
			default:
				close(targetClient.Send)
				h.Clients.Delete(targetClient.ID)
			}
		}
	}
}
