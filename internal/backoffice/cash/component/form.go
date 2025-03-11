package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(cash *cash.Cash, errors validation.Errors) Node {
	return Form(
		If(cash.ID == 0, htmx.Post(path.BackofficeCashes())),
		If(cash.ID != 0, htmx.Patch(path.BackofficeCash(cash.ID))),
		c.Input("Назва", "name", cash.Name, nil, errors.Get("name")),
		c.Input("Сума", "formula", cash.Formula, nil, errors.Get("sum")),
		Div(Class("d-flex justify-content-between"),
			c.Submit(cash.ID),
			If(cash.ID != 0,
				Div(
					Div(Class("btn btn-outline-danger"),
						htmx.Delete(path.BackofficeCash(cash.ID)),
						htmx.Confirm("Are you sure?"),
						Text("Видалити"),
					),
				),
			),
		),
	)
}
