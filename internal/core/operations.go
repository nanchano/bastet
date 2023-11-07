package core

import "context"

// CreateEvent coantains the business logic to create new events on the service,
// calling the repository to actually store the new event.
func (s service) CreateEvent(ctx context.Context, e Event) (*Event, error) {
	if err := e.Validate(); err != nil {
		s.logger.Error("Validation error", "error", err)
		return nil, err
	}

	event, err := s.repo.CreateEvent(ctx, e)
	if err != nil {
		error := NewError("Failed creating event", ErrorUnknown)
		s.logger.Error(error.msg, "error", error)
		return nil, error
	}

	return event, nil
}

// GetEvent contains the business logic to read event details,
// calling the repository to get the relevant payload.
func (s service) GetEvent(ctx context.Context, id int64) (*Event, error) {
	if id == 0 {
		error := NewError("Invalid event ID: 0", ErrorInvalidArgument)
		s.logger.Error(error.msg, "error", error)
		return nil, error
	}

	event, err := s.repo.GetEvent(ctx, id)
	if err != nil {
		error := NewError("Event not found", ErrorNotFound)
		s.logger.Error(error.msg, "error", error)
		return nil, error
	}

	return event, nil
}

// UpdateEvent contains the business logic to update events,
// from the service, calling the repository to store the updated event.
func (s service) UpdateEvent(ctx context.Context, e Event) (*Event, error) {
	if err := e.Validate(); err != nil {
		s.logger.Error("Validation error", "error", err)
		return nil, err
	}

	event, err := s.repo.UpdateEvent(ctx, e)
	if err != nil {
		error := NewError("Failed Updating event", ErrorUnknown)
		s.logger.Error(error.msg, "error", error)
		return nil, error
	}

	return event, nil
}

// DeleteEvent contains the business logic for the deletion of the event
// from the service, calling the repository to actually delete it.
func (s service) DeleteEvent(ctx context.Context, id int64) error {
	if id == 0 {
		error := NewError("Invalid event ID: 0", ErrorInvalidArgument)
		s.logger.Error(error.msg, "error", error)
		return error
	}

	err := s.repo.DeleteEvent(ctx, id)
	if err != nil {
		error := NewError("Event not found", ErrorNotFound)
		s.logger.Error(error.msg, "error", error)
		return error
	}

	return nil
}
