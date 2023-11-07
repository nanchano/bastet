package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Ping checks the server is running
func (s server) Ping(w http.ResponseWriter, r *http.Request) {
	renderResponse(w, "pong")
}

// Create event receives a request and calls the service to create a new event
// given the payload, sending the event itself as a response afterwards.
// If unsuccessful, an error response will be sent.
func (s server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		s.logger.Error("Failed parsing the request")
		renderErrorResponse(w, r, err)
		return
	}

	e, err := s.service.CreateEvent(r.Context(), event.toCoreEvent())
	if err != nil {
		s.logger.Error("Failed creating the event")
		renderErrorResponse(w, r, err)
		return
	}

	renderResponse(w, e)
}

// GetEvent receives a request and calls the service for the required
// event given the ID, return the event as a response.
// If unsuccessful, an error response will be sent.
func (s server) GetEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "event_id"), 10, 64)
	if err != nil {
		s.logger.Error("Failed parsing the request, could not parse the event ID")
		renderErrorResponse(w, r, err)
		return
	}

	e, err := s.service.GetEvent(r.Context(), id)
	if err != nil {
		s.logger.Error("Failed finding the event")
		renderErrorResponse(w, r, err)
		return
	}

	renderResponse(w, e)
}

// UpdateEvent receives a request and calls the service to
// update an event, sending a response with the updated event afterwards.
// If unsuccessful, an error response will be sent.
func (s server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		s.logger.Error("Failed parsing the request")
		renderErrorResponse(w, r, err)
		return
	}

	if event.ID == 0 {
		id, err := strconv.ParseInt(chi.URLParam(r, "event_id"), 10, 64)
		if err != nil {
			s.logger.Error("Event ID not found")
			renderErrorResponse(w, r, err)
			return
		}
		event.ID = id
	}

	e, err := s.service.UpdateEvent(r.Context(), event.toCoreEvent())
	if err != nil {
		s.logger.Error("Failed updating the event")
		renderErrorResponse(w, r, err)
		return
	}

	renderResponse(w, e)
}

// DeleteEvent receives a request and calls the service to
// delete an event by ID, sending an empty response afterwards.
// If unsuccessful, an error response will be sent.
func (s server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "event_id"), 10, 64)
	if err != nil {
		s.logger.Error("Failed parsing the request, could not parse the event ID")
		renderErrorResponse(w, r, err)
		return
	}

	err = s.service.DeleteEvent(r.Context(), id)
	if err != nil {
		s.logger.Error("Failed deleting the event")
		renderErrorResponse(w, r, err)
		return
	}

	renderResponse(w, map[string]interface{}{})
}
