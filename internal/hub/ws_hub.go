package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/InstaySystem/is-be/internal/service"
	"github.com/InstaySystem/is-be/internal/types"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WSClient struct {
	Hub          *WSHub
	Conn         *websocket.Conn
	Send         chan []byte
	ID           string
	ClientID     int64
	Type         string
	DepartmentID *int64
	ActiveChats  map[int64]bool
}

func NewWSClient(hub *WSHub, conn *websocket.Conn, clientID int64, clientType string, departmentID *int64) *WSClient {
	return &WSClient{
		hub,
		conn,
		make(chan []byte, 256),
		uuid.NewString(),
		clientID,
		clientType,
		departmentID,
		make(map[int64]bool),
	}
}

type WSHub struct {
	Clients     map[string]map[string]*WSClient
	Register    chan *WSClient
	Unregister  chan *WSClient
	SendMessage chan *MessagePayload
	ChatSvc     service.ChatService
}

type MessagePayload struct {
	TargetKey string
	Data      []byte
}

func NewWSHub(chatSvc service.ChatService) *WSHub {
	return &WSHub{
		make(map[string]map[string]*WSClient),
		make(chan *WSClient),
		make(chan *WSClient),
		make(chan *MessagePayload),
		chatSvc,
	}
}

func (c *WSClient) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, data, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("%v", err)
			}
			break
		}

		var msg types.CreateMessageRequest
		if err = json.Unmarshal(data, &msg); err != nil {
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		message, err := c.Hub.ChatSvc.CreateMessage(ctx, c.ClientID, c.DepartmentID, c.Type, msg)
		if err != nil {
			break
		}

		messageBytes, _ := json.Marshal(message)

		var targetKey string
		if c.Type == "guest" {
			targetKey = fmt.Sprintf("dept_%d", message.Chat.DepartmentID)
		} else {
			targetKey = fmt.Sprintf("guest_%d", message.Chat.OrderRoomID)
		}

		targetPayload := &MessagePayload{
			TargetKey: targetKey,
			Data:      messageBytes,
		}
		c.Hub.SendMessage <- targetPayload
	}
}

func (c *WSClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.Register:
			key := client.getKey()
			if _, ok := h.Clients[key]; !ok {
				h.Clients[key] = make(map[string]*WSClient)
			}
			h.Clients[key][client.ID] = client

		case client := <-h.Unregister:
			key := client.getKey()
			if conns, ok := h.Clients[key]; ok {
				if _, exists := conns[client.ID]; exists {
					delete(conns, client.ID)
					close(client.Send)
					if len(conns) == 0 {
						delete(h.Clients, key)
					}
				}
			}

		case msg := <-h.SendMessage:
			if conns, ok := h.Clients[msg.TargetKey]; ok {
				for _, client := range conns {
					select {
					case client.Send <- msg.Data:
					default:
						close(client.Send)
						delete(conns, client.ID)
					}
				}

				if len(conns) == 0 {
					delete(h.Clients, msg.TargetKey)
				}
			}
		}
	}
}

func (c *WSClient) getKey() string {
	if c.Type == "guest" {
		return fmt.Sprintf("guest_%d", c.ClientID)
	}
	return fmt.Sprintf("dept_%d", *c.DepartmentID)
}
