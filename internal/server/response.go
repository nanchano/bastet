package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nanchano/bastet/internal/core"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error core.Error `json:"error"`
}

// renderErrorResponse parses the error type and writes an HTTP error response
func renderErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	resp := ErrorResponse{
		Error: core.NewError("Internal Server Error", status),
	}

	var ierr *core.Error
	if !errors.As(err, &ierr) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		switch ierr.Code() {
		case core.ErrorNotFound:
			status = http.StatusNotFound
		case core.ErrorInvalidArgument:
			status = http.StatusBadRequest
		case core.ErrorUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	}

	renderResponse(w, resp)
}

// renderResponse sends a JSON response given the payload.
func renderResponse(w http.ResponseWriter, v interface{}) {
	payload, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
