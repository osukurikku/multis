package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// WebSocketURL is the Ripple API endpoint that accepts websocket connections.
const WebSocketURL = "wss://api.ripple.moe/api/v1/ws"

// ListenToWS connects to the Ripple WebSocket API and listens to new messages.
// Once a new message is received, it forwards it to the correct function
// on the typeMap.
func ListenToWS(typeMap map[string]func(Message, *websocket.Conn)) error {
	conn, _, err := (&websocket.Dialer{}).Dial(WebSocketURL, nil)
	if err != nil {
		return err
	}
	for {
		var m Message
		err := conn.ReadJSON(&m)
		if err != nil {
			conn.Close()
			return err
		}
		f := typeMap[m.Type]
		if f != nil {
			go f(m, conn)
		}
	}
}

// A Message received from the Ripple API
type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data,omitempty"`
}

// Unmarshal parses the JSON-encoded data in m.Data and stores the result into
// v.
func (m Message) Unmarshal(v interface{}) error {
	return json.Unmarshal(m.Data, v)
}
