package common

import (
	"fmt"
	"strings"
	"time"

	"github.com/tksasha/balance/pkg/validation"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

type BaseComponent struct{}

func NewBaseComponent() *BaseComponent {
	return &BaseComponent{}
}

func (c *BaseComponent) Date(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(time.DateOnly)
}

func (c *BaseComponent) Errors(attribute string, errors validation.Errors) Node {
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

func (c *BaseComponent) Money(sum float64) string {
	return message.NewPrinter(language.Ukrainian).Sprintf("%0.2f", sum)
}
