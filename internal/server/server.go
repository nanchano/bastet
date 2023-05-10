package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nanchano/bastet/internal/core"
	"golang.org/x/exp/slog"
)

// server represents the HTTP layer of the app, requiring a server to act.
type server struct {
	logger  *slog.Logger
	service core.BastetService
}

// New starts a new server given the service.
func New(l *slog.Logger, s core.BastetService) *server {
	return &server{
		logger:  l,
		service: s,
	}
}

// Start starts the server
func (s server) Start() {
	s.logger.Info("Starting server on port 3333")
	r := chi.NewRouter()

	r.Use(middleware.Logger, middleware.RequestID)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/bastet", s.CreateEvent)
	r.Get("/bastet/{event_id}", s.GetEvent)
	r.Put("/bastet/{event_id}", s.UpdateEvent)
	r.Delete("/bastet/{event_id}", s.DeleteEvent)

	http.ListenAndServe(":3333", r)

	s.logger.Info("Serving")
}
