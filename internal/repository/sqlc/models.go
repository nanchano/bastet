// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"time"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Category    string
	Location    string
	Publisher   string
	Lineup      []string
	StartTS     time.Time
	EndTS       time.Time
	CreatedAt   time.Time
}
