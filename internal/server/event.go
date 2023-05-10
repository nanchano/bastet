package server

import (
	"time"

	"github.com/nanchano/bastet/internal/core"
)

// Event defines model for Event.
type Event struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Location    string    `json:"location"`
	Publisher   string    `json:"publisher"`
	Lineup      []string  `json:"lineup"`
	StartTS     time.Time `json:"start_ts"`
	EndTS       time.Time `json:"end_ts"`
}

// toCoreEvent transforms the server event into a core event
func (e Event) toCoreEvent() core.Event {
	return core.Event{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Category:    e.Category,
		Location:    e.Location,
		Publisher:   e.Publisher,
		Lineup:      e.Lineup,
		StartTS:     e.StartTS,
		EndTS:       e.EndTS,
	}
}
