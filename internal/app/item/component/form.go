package component

import (
	"time"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/validator"
	. "maragu.dev/gomponents" //nolint:stylecheck
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Form(item *item.Item, errors validator.Errors) Node {
	return Form(
		c.Date(item.Date, errors.Get("date")),
	)
}

func (c *Component) Date(date time.Time, message *string) Node {
	classes := components.Classes{
		"form-control": true,
		"is-invalid":   message != nil,
	}

	return Div(Class("mb-3"),
		Label(Class("form-label"), Text("Дата")),
		Input(classes, Value(date.Format(time.DateOnly))),
		Iff(message != nil, func() Node { return Div(Class("invalid-feedback"), Text(*message)) }),
	)
}
