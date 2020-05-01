package models

import "time"

// Game represents a game (a single playthrough of a beatmap) of a match
type Game struct {
	ID           int    `gorm:"primary_key"`
	Match        Match  `json:"-"`
	MatchID      int64  `json:",string"`
	HostID       int    `json:",int"`
	HostUserName string `json:",string"`
	Name         string
	BeatmapID    int
	Mods         int64
	GameMode     int
	Scores       string `gorm:"type:text"`
	CreatedAt    time.Time
}
