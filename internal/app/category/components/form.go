package components

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *CategoryComponent) form(category *category.Category, errors validation.Errors) Node {
	return Form(
		Div(
			Class("mb-3"),
			Label(Text("name")),
			Input(Value(category.Name)),
			c.Errors("name", errors),
		),
	)
}
