package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(cash *cash.Cash, errors validation.Errors) Node {
	return Form(
		If(cash.ID == 0, htmx.Post(paths.CreateBackofficeCash())),
		If(cash.ID != 0, htmx.Patch(paths.UpdateBackofficeCash(cash.ID))),
		Div(Class("mb-3"),
			Label(Class("form-label"), Text("Валюта")),
			Select(Class("form-select"), Name("currency"),
				c.CurrencyOptions(currency.GetCode(cash.Currency))),
		),
		c.Input("Назва", "name", cash.Name, nil, errors.Get("name"), AutoFocus()),
		c.Input("Сума", "formula", cash.Formula, nil, errors.Get("sum")),
		c.Input("Група", "supercategory", strconv.Itoa(cash.Supercategory), nil, errors.Get("supercategory")),
		Div(Class("d-flex justify-content-between"),
			c.Submit(cash.ID),
			If(cash.ID != 0,
				Button(Class("btn btn-outline-danger"),
					htmx.Delete(paths.DeleteBackofficeCash(cash.ID)),
					htmx.Confirm("Are you sure?"),
					Text("Видалити"),
				),
			),
		),
	)
}
