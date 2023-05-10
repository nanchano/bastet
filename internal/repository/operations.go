package repository

import (
	"context"
	"fmt"

	"github.com/nanchano/bastet/internal/core"
	"github.com/nanchano/bastet/internal/repository/sqlc"
)

// CreateEvent creates a new event in the repository given the payload.
func (r *repository) CreateEvent(ctx context.Context, e core.Event) (*core.Event, error) {
	params := sqlc.CreateEventParams{
		Name:        e.Name,
		Description: e.Description,
		Category:    e.Category,
		Location:    e.Location,
		Publisher:   e.Publisher,
		Lineup:      e.Lineup,
		StartTS:     e.StartTS,
		EndTS:       e.EndTS,
	}

	event, err := r.querier.CreateEvent(ctx, params)
	if err != nil {
		return nil, err
	}

	return toCoreEvent(event), nil
}

// GetEvent retrieves an event from the repository given an ID.
func (r *repository) GetEvent(ctx context.Context, id int64) (*core.Event, error) {
	e, err := r.querier.GetEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	return toCoreEvent(e), nil
}

// UpdateEvent updates an event from the repository given the payload.
func (r *repository) UpdateEvent(ctx context.Context, e core.Event) (*core.Event, error) {
	params := sqlc.UpdateEventParams{
		Name:        toNullString(e.Name),
		Description: toNullString(e.Description),
		Category:    toNullString(e.Category),
		Location:    toNullString(e.Location),
		Publisher:   toNullString(e.Publisher),
		Lineup:      e.Lineup,
		StartTS:     toNullTime(e.StartTS),
		EndTS:       toNullTime(e.EndTS),
		ID:          e.ID,
	}
	fmt.Printf("%v", params)

	event, err := r.querier.UpdateEvent(ctx, params)
	if err != nil {
		return nil, err
	}

	return toCoreEvent(event), nil
}

// DeleteEvent deletes an event from the repository given an ID.
func (r *repository) DeleteEvent(ctx context.Context, id int64) error {
	err := r.querier.DeleteEvent(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
