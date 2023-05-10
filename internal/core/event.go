package core

import "time"

// Event represents the main use case of the app, an event to occurr
// In a given timeframe and location, with a lineup and description.
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
}

// Validate checks whether the event is valid by testing against some of its attributes.
func (e Event) Validate() error {
	if e.ID < 0 {
		return NewError("Event IDs must be positive", ErrorValidationFailed)
	}

	if e.EndTS.Before(e.StartTS) {
		return NewError("The event can't start after it has ended", ErrorValidationFailed)
	}

	return nil
}
