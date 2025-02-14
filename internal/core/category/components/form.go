package components

import (
	"github.com/tksasha/balance/internal/core/category"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func form(category *category.Category) Node {
	return Form(
		Div(
			Class("mb-3"),
			Label(Text("name")),
			Input(Value(category.Name)),
		),
	)
}
