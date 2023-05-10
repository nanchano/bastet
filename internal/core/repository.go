package core

import "context"

// Repository defines the methods needed to become a repo in the Bastet service.
// Implementations need to be able to Create, Read, Update and Delete events
type Repository interface {
	CreateEvent(ctx context.Context, e Event) (*Event, error)
	GetEvent(ctx context.Context, id int64) (*Event, error)
	UpdateEvent(ctx context.Context, e Event) (*Event, error)
	DeleteEvent(ctx context.Context, id int64) error
}
