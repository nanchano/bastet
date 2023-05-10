package core

const (
	ErrorUnknown int = iota
	ErrorNotFound
	ErrorInvalidArgument
	ErrorValidationFailed
)

// Error represents an error ocurred on the service layer due to business logic failing.
type Error struct {
	msg  string
	code int
}

// NewError creates a new Error given a message and a code.
func NewError(msg string, code int) Error {
	return Error{
		msg:  msg,
		code: code,
	}
}

// Error returns the message of the error.
func (e Error) Error() string {
	return e.msg
}

// Code returns the code representing this error.
func (e *Error) Code() int {
	return e.code
}
