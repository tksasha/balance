package errors

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

type NotFoundError struct {
	message string
}

func NewNotFoundError(args ...string) error {
	message := "not found"

	if len(args) == 1 {
		message = args[0]
	}

	return NotFoundError{
		message: message,
	}
}

func (e NotFoundError) Error() string {
	return e.message
}
