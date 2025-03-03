package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) form(category *category.Category, errors validation.Errors) Node {
	_ = errors

	return Form(
		Div(
			Class("mb-3"),
			Label(Text("name")),
			Input(Value(category.Name)),
		),
	)
}
