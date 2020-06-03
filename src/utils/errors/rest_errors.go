package errors

import (
	"fmt"
	"net/http"
)

// RestError represents a rest api error.
type RestError interface {
	// Error return the error description.
	Error() string
	// Status returns the http status code related with the error.
	Status() int
	// Message returns a readeble message about this error.
	Message() string
	// Causes returns an array of causes behind this error.
	Causes() []interface{}
}

type restError struct {
	error   string        `json:"error"`
	status  int           `json:"status"`
	message string        `json:"message"`
	causes  []interface{} `json:"causes"`
}

func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v", e.message, e.status, e.error, e.causes)
}

func (e restError) Status() int {
	return e.status
}

func (e restError) Message() string {
	return e.message
}

func (e restError) Causes() []interface{} {
	return e.causes
}

func NewRestError(message string, status int, error string, causes []interface{}) RestError {
	return restError{
		message: message,
		status:  status,
		error:   error,
		causes:  causes,
	}
}

func NewBadRequestError(message string) RestError {
	return restError{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}
