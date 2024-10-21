package handlers

import (
	"fmt"
)

type FormParsingError struct {
	message string
}

func NewFormParsingError(err error) error {
	return &FormParsingError{
		message: fmt.Sprintf("form parsing error: %v", err),
	}
}

func (e *FormParsingError) Error() string {
	return e.message
}
