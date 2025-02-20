package component

import (
	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(cash *cash.Cash, errors validation.Errors) Node {
	return Form(
		Div(
			Class("mb-3"),
			Label(
				For("cash_name"),
				Class("form-label"),
				Text("Name"),
			),
			Input(
				Class("form-control"),
				ID("cash_name"),
				Value(cash.Name),
			),
			components.Errors("name", errors),
		),
		Div(
			Class("mb-3"),
			Label(
				For("cash_sum"),
				Class("form-label"),
				Text("Sum"),
			),
			Input(
				Class("form-control"),
				ID("cash_id"),
				Value(cash.Formula),
			),
		),
		Button(
			Type("submit"),
			Class("btn btn-primary"),
			If(cash.ID == 0, Text("Create")),
			If(cash.ID != 0, Text("Update")),
		),
	)
}
