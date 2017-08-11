package webserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"zxq.co/ripple/vinse/vinstore/models"
)

func atoiMin(s string, def int) int {
	i, _ := strconv.Atoi(s)
	if i <= 0 {
		return def
	}
	return i
}

func limit(i, lim int) int {
	if i > lim {
		return lim
	}
	return i
}

// Matches handles requests for /api/matches
func Matches(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var matches []models.Match

	query := r.URL.Query()

	l := limit(atoiMin(query.Get("l"), 100), 1000)
	q := DB.Limit(l).Offset(l * atoiMin(query.Get("p"), 0))
	if val, ok := query["id"]; ok {
		q = q.Where("id IN (?)", val)
	}
	q.Order("created_at DESC").Find(&matches)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(matches)
}

type jsonScore struct {
	Score  int64 `json:"score"`
	Mods   int64 `json:"mods"`
	Passed bool  `json:"passed"`
	Failed bool  `json:"failed"`
}

type jsonScoreCamel struct {
	Score  int64
	Mods   int64
	Passed bool
	Failed bool
}

func (j *jsonScoreCamel) UnmarshalJSON(data []byte) error {
	var jScore jsonScore
	err := json.Unmarshal(data, &jScore)
	if err != nil {
		return err
	}
	// conversion between two structs with same fields but different field tags
	// is legal since Go 1.8
	*j = jsonScoreCamel(jScore)
	return nil
}

type jsonGame struct {
	models.Game
	Scores map[string]jsonScoreCamel
}

// gamesForJSON basically changes the `Score` field from being a string
// containing a JSON object to being an actual JSON object.
func gamesForJSON(games []models.Game) []jsonGame {
	jsonGames := make([]jsonGame, len(games))
	for i, game := range games {
		jsonGames[i].Game = game
		jScoreMap := make(map[string]jsonScoreCamel, 16)
		json.Unmarshal([]byte(game.Scores), &jScoreMap)
		jsonGames[i].Scores = jScoreMap
	}
	return jsonGames
}

// Games handles requests for /api/games
func Games(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var games []models.Game

	query := r.URL.Query()

	l := limit(atoiMin(query.Get("l"), 100), 1000)
	q := DB.Limit(l).Offset(l * atoiMin(query.Get("p"), 0))

	if val, ok := query["id"]; ok {
		q = q.Where("id IN (?)", val)
	}
	if val, ok := query["match_id"]; ok {
		q = q.Where("match_id IN (?)", val)
	}

	q.Order("created_at DESC").Find(&games)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(gamesForJSON(games))
}
