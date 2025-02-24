package component

import (
	"fmt"
	"strings"

	"github.com/tksasha/validator"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Errors(attribute string, errors validator.Errors) Node {
	if !has(attribute, errors) {
		return nil
	}

	return Div(
		Class("invalid-feedback"),
		Text(get(attribute, errors)),
	)
}

func has(attribute string, errors validator.Errors) bool {
	if errors == nil {
		return false
	}

	return errors.Has(attribute)
}

func get(attribute string, errors validator.Errors) string {
	if errors == nil {
		return ""
	}

	messages := strings.Join(errors.Get(attribute), ", ")

	return fmt.Sprintf("%s: %s", attribute, messages)
}
