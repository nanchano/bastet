package core

import "context"

// CreateEvent coantains the business logic to create new events on the service,
// calling the repository to actually store the new event.
func (s *service) CreateEvent(ctx context.Context, e Event) (*Event, error) {
	s.logger.Info("Creating event with the following parameters: %v", e)

	if error := e.Validate(); error != nil {
		return nil, error
	}

	event, err := s.repo.CreateEvent(ctx, e)
	if err != nil {
		error := NewError("Failed creating event: %v", ErrorUnknown)
		s.logger.Error(error.msg, error)
		return nil, error
	}

	s.logger.Info("Succesfully created the event")
	return event, nil
}

// GetEvent contains the business logic to read event details,
// calling the repository to get the relevant payload.
func (s *service) GetEvent(ctx context.Context, id int64) (*Event, error) {
	s.logger.Info("Getting event with ID: %s", id)
	if id == 0 {
		error := NewError("Invalid event ID: 0", ErrorInvalidArgument)
		s.logger.Error(error.msg, error)
		return nil, error
	}

	event, err := s.repo.GetEvent(ctx, id)
	if err != nil {
		error := NewError("Event not found", ErrorNotFound)
		s.logger.Error(error.msg, error)
		return nil, error
	}

	s.logger.Info("Succesfully found the event: %v", event)
	return event, nil
}

// UpdateEvent contains the business logic to update events,
// from the service, calling the repository to store the updated event.
func (s *service) UpdateEvent(ctx context.Context, e Event) (*Event, error) {
	s.logger.Info("Updating event with the following parameters: %v", e)
	if error := e.Validate(); error != nil {
		return nil, error
	}

	event, err := s.repo.UpdateEvent(ctx, e)
	if err != nil {
		error := NewError("Failed Updating event: %v", ErrorUnknown)
		s.logger.Error(error.msg, error)
		return nil, error
	}

	s.logger.Info("Succesfully updated the event")
	return event, nil
}

// DeleteEvent contains the business logic for the deletion of the event
// from the service, calling the repository to actually delete it.
func (s *service) DeleteEvent(ctx context.Context, id int64) error {
	s.logger.Info("Deleting event with ID: %s", id)
	if id == 0 {
		error := NewError("Invalid event ID: 0", ErrorInvalidArgument)
		s.logger.Error(error.msg, error)
		return error
	}

	err := s.repo.DeleteEvent(ctx, id)
	if err != nil {
		error := NewError("Event not found", ErrorNotFound)
		s.logger.Error(error.msg, error)
		return error
	}

	s.logger.Info("Succesfully deleted the event")
	return nil
}
