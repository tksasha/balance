package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validation"
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
