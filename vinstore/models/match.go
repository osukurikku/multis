package models

import "time"

// Match represents a multiplayer match on Ripple.
type Match struct {
	ID        int64 `gorm:"primary_key" json:",string"`
	Name      string
	CreatedAt time.Time
}

// GenerateMatchID creates the ID of the match.
// The reason for this is that the IDs given by pep.py, and thus the Ripple API,
// can repeat themselves. This is due to the fact that on a restart of pep.py,
// the IDs start from one again, and so to overcome this two games are
// considered of the same match only if they share the same match ID and have
// been created at most 29 minutes and 59 seconds apart (in a "worst case
// scenario")
func GenerateMatchID(id int, refTime time.Time) int64 {
	return (refTime.Unix()/(60*15))<<32 | int64(id)
}
