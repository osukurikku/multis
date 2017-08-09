package models

// MatchRedirect contains what matches some match IDs are referring to.
// ... It's complicated.
type MatchRedirect struct {
	ID        int64 `gorm:"primary_key"`
	CanonicID int64
}
