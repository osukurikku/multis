package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/osukurikku/multis/vinstore/models"
	"github.com/osukurikku/multis/vinstore/webserver"
)

// Client specifies how to handle the messages received from the Ripple
// WebSocket API
type Client struct {
	DB *gorm.DB
}

// Connected is a function to be fired once a message of type "connected"
// is received.
func (c *Client) Connected(_ Message, conn *websocket.Conn) {
	fmt.Println("Connected")
	conn.WriteJSON(Message{
		Type: "subscribe_mp_complete_match",
	})
	go c.Pinger(conn)
}

// Pinger sends a message of type "ping" every 10 seconds to the conn.
func (c *Client) Pinger(conn *websocket.Conn) {
	for {
		err := conn.WriteJSON(Message{
			Type: "ping",
		})
		if err != nil {
			return
		}
		time.Sleep(10 * time.Second)
	}
}

// WSGame is a multiplayer game that has just been finished, as represented by
// the Ripple API.
type WSGame struct {
	BeatmapID    int             `json:"beatmap_id"`
	GameMode     int             `json:"game_mode"`
	ID           int             `json:"id"`
	Mods         int64           `json:"mods"`
	Name         string          `json:"name"`
	HostID       int             `json:"host_id"`
	HostUserName string          `json:"host_user_name"`
	Scores       json.RawMessage `json:"scores"`
}

// NewCompletedMatch handles messages of type new_completed_match.
func (c *Client) NewCompletedMatch(m Message, _ *websocket.Conn) {
	var game WSGame
	err := m.Unmarshal(&game)
	if err != nil {
		fmt.Println(err)
		return
	}
	if game.ID == 0 {
		return
	}

	now := time.Now()

	matchID := models.GenerateMatchID(game.ID, now)
	fmt.Printf("New game received, id: %d (%d)\n", game.ID, matchID)

	oldMatchID := models.GenerateMatchID(game.ID, now.Add(-time.Minute*15))

	// Check if there's already a match redirect for this matchID
	mr := models.MatchRedirect{}
	c.DB.Where("id IN (?)", []int64{matchID, oldMatchID}).Order("id DESC").First(&mr)

	if mr.ID == 0 {
		// Match doesn't exist, we need to create it
		c.DB.Create(&models.Match{
			ID:           matchID,
			Name:         game.Name,
			HostID:       game.HostID,
			HostUserName: game.HostUserName,
		})
		c.DB.Create(&models.MatchRedirect{
			ID:        matchID,
			CanonicID: matchID,
		})
		mr.CanonicID = matchID
	} else if mr.ID != matchID {
		// Latest match redirect doesn't exist, so we need to create it
		c.DB.Create(&models.MatchRedirect{
			ID:        matchID,
			CanonicID: mr.CanonicID,
		})
	}

	createdGame := models.Game{
		MatchID:      mr.CanonicID,
		Name:         game.Name,
		BeatmapID:    game.BeatmapID,
		Mods:         game.Mods,
		GameMode:     game.GameMode,
		HostID:       game.HostID,
		HostUserName: game.HostUserName,
		Scores:       string(game.Scores),
	}
	c.DB.Create(&createdGame)

	webserver.Stream(createdGame)
}

// GenerateTypeMap creates a typeMap for use with ListenToWS.
func (c *Client) GenerateTypeMap() map[string]func(Message, *websocket.Conn) {
	return map[string]func(Message, *websocket.Conn){
		"connected":           c.Connected,
		"new_completed_match": c.NewCompletedMatch,
	}
}
