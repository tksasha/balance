package validation

import (
	"fmt"
	"strings"

	utilsstrings "github.com/tksasha/utils/strings"
)

type Errors map[string][]string //nolint:errname

func (e Errors) Get(attribute string) []string {
	return e[attribute]
}

func (e Errors) Has(attribute string) bool {
	return len(e[attribute]) > 0
}

func (e Errors) Error() string {
	errors := []string{}

	for attribute, messages := range e {
		errors = append(errors, fmt.Sprintf("%s: %s", attribute, strings.Join(messages, ", ")))
	}

	return strings.Join(errors, "; ")
}

func (e Errors) add(attribute, message string) {
	attribute = utilsstrings.ToSnakeCase(attribute)

	e[attribute] = append(e[attribute], message)
}

func (e Errors) exists() bool {
	return len(e) != 0
}
