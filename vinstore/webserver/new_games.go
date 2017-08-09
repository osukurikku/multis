package webserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/julienschmidt/sse"
	"zxq.co/ripple/vinse/vinstore/models"
)

var streamer = sse.New()

func init() {
	go func() {
		for {
			streamer.SendBytes("", "", []byte{'h'})
			time.Sleep(time.Second * 5)
		}
	}()
}

// Stream sends a new multiplayer game on the streamer
func Stream(game models.Game) {
	scores := make(map[string]jsonScoreCamel)
	json.Unmarshal([]byte(game.Scores), &scores)
	jGame := jsonGame{
		Game:   game,
		Scores: scores,
	}

	streamer.SendJSON("", "", jGame)
}

// NewGames handle the HTTP request for new_games
func NewGames(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	streamer.ServeHTTP(w, nil)
}
