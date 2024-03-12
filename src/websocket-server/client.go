package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	hub *Hub

	conn *websocket.Conn

	send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.WithError(err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var messageJson Command
		err = json.Unmarshal([]byte(message), &messageJson)
		if err != nil {
			log.WithField("Message", message).WithError(err)
			continue
		}
		switch messageJson.CommandCode {
		case Host:
			data := &hubData{
				Client:   c,
				Nickname: messageJson.Nickname,
			}
			c.hub.create <- data
		case Start:
			data := &hubData{
				Client: c,
				GameId: messageJson.GameId,
			}
			c.hub.start <- data
		case Join:
			data := &hubData{
				Client:   c,
				GameId:   messageJson.GameId,
				Nickname: messageJson.Nickname,
			}
			c.hub.join <- data
		case Kick:
			data := &hubData{
				Client:   c,
				GameId:   messageJson.GameId,
				Nickname: messageJson.Nickname,
			}
			c.hub.kick <- data
		case Play:
			socketData := &hubData{
				Client: c,
				GameId: messageJson.GameId,
			}
			data := &gameData{
				hubData: *socketData,
				Command: messageJson.GameData,
			}
			c.hub.play <- data
		case Leave:
			socketData := &hubData{
				Client: c,
				GameId: messageJson.GameId,
			}
			c.hub.leave <- socketData
		case Status:
			socketData := &hubData{
				Client: c,
				GameId: messageJson.GameId,
			}
			c.hub.status <- socketData
		}
	}
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
