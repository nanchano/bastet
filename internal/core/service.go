package core

import (
	"context"

	"golang.org/x/exp/slog"
)

// BastetService is the main usecase of the app, centered around event CRUD operations.
type BastetService interface {
	CreateEvent(ctx context.Context, e Event) (*Event, error)
	GetEvent(ctx context.Context, id int64) (*Event, error)
	UpdateEvent(ctx context.Context, e Event) (*Event, error)
	DeleteEvent(ctx context.Context, id int64) error
}

// Service implements a BastetService, and will contain the operations and business logic
// necessary while interacting with our 3rd party adapters.
type service struct {
	logger *slog.Logger
	repo   Repository
}

// NewService creates a new BastetService given a logger anr repository
func NewService(logger *slog.Logger, repo Repository) *service {
	return &service{
		logger: logger,
		repo:   repo,
	}
}
