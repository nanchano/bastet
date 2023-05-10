package repository

import (
	"database/sql"
	"time"

	"github.com/nanchano/bastet/internal/core"
	"github.com/nanchano/bastet/internal/repository/sqlc"
)

// toCoreEvent takes a model from the SQL layer, the sqlc.Event, and converts it into a core.Event
func toCoreEvent(e sqlc.Event) *core.Event {
	return &core.Event{
		ID:          e.ID,
		Description: e.Description,
		Name:        e.Name,
		Category:    e.Category,
		Location:    e.Location,
		Publisher:   e.Publisher,
		Lineup:      e.Lineup,
		StartTS:     e.StartTS,
		EndTS:       e.EndTS,
	}
}

// toNullString converts a string into a sql.NullString
// It's valid if the string is not ""
func toNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// toNullTime converts a time.Time object into sql.NullTime.
// It's a valid Null time if Time is not zero and is not equal to the unix timestamp.
func toNullTime(t time.Time) sql.NullTime {
	unixStart := "1970-01-01 00:00:00 +0000"
	layout := "2006-01-02 03:04:05 -0700"
	unix, _ := time.Parse(layout, unixStart)
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero() && !t.Equal(unix),
	}
}
