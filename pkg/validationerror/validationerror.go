package validationerror

import (
	"fmt"
	"strings"

	utilsstrings "github.com/tksasha/utils/strings"
)

const (
	IsRequired    = "is required"
	IsInvalid     = "is invalid"
	AlreadyExists = "already exists"
	InternalError = "internal error"
)

type ValidationError map[string][]string

func New() ValidationError {
	return ValidationError{}
}

func NewWithMessage(attribute, message string) ValidationError {
	err := ValidationError{}

	err.Set(attribute, message)

	return err
}

func (e ValidationError) Error() string {
	result := []string{}

	for attribute, messages := range e {
		result = append(result, fmt.Sprintf("%s: %s", attribute, strings.Join(messages, ", ")))
	}

	return strings.Join(result, "; ")
}

func (e ValidationError) Set(attribute, message string) {
	attribute = utilsstrings.ToSnakeCase(attribute)

	e[attribute] = append(e[attribute], message)
}

func (e ValidationError) Get(attribute string) []string {
	return e[attribute]
}

func (e ValidationError) Exists() bool {
	return len(e) != 0
}

func (e ValidationError) Has(attribute string) bool {
	return len(e[attribute]) > 0
}
