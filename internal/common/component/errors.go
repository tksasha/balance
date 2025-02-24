package component

import (
	"fmt"
	"strings"

	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Errors(attribute string, errors validation.Errors) Node {
	if !has(attribute, errors) {
		return nil
	}

	return Div(
		Class("invalid-feedback"),
		Text(get(attribute, errors)),
	)
}

func has(attribute string, errors validation.Errors) bool {
	if errors == nil {
		return false
	}

	return errors.Has(attribute)
}

func get(attribute string, errors validation.Errors) string {
	if errors == nil {
		return ""
	}

	messages := strings.Join(errors.Get(attribute), ", ")

	return fmt.Sprintf("%s: %s", attribute, messages)
}
