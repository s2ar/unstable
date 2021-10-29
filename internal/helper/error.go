package helper

import (
	"fmt"
)

const (
	defaultInternalServerErrorMessage = "internal server error"
)

type NotFoundErr struct {
	message string
}

func (n NotFoundErr) Error() string {
	return n.message
}

func NewNotFoundErr(template string, args ...interface{}) error {
	return NotFoundErr{
		message: fmt.Sprintf(template, args...), // nolint: vet
	}
}

type HTTPError struct {
	Status  int
	Err     error
	message string
}

func (h *HTTPError) Error() string {
	if h.message != "" {
		return h.message
	}

	return h.Err.Error()
}

func NewHTTPError(err error, status int, format string, args ...interface{}) *HTTPError {
	return &HTTPError{
		Status:  status,
		Err:     err,
		message: fmt.Sprintf(format, args...), // nolint: vet
	}
}
