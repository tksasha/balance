package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(cash *cash.Cash, errors validation.Errors) Node {
	return Form(
		If(cash.ID == 0, htmx.Post(path.CreateBackofficeCash(nil))),
		If(cash.ID != 0, htmx.Patch(path.UpdateBackofficeCash(nil, cash.ID))),
		Div(Class("mb-3"),
			Label(Class("form-label"), Text("Валюта")),
			Select(Class("form-select"), Name("currency"),
				c.currencyOptions(currency.GetCode(cash.Currency))),
		),
		c.Input("Назва", "name", cash.Name, nil, errors.Get("name")),
		c.Input("Сума", "formula", cash.Formula, nil, errors.Get("sum")),
		c.Input("Група", "supercategory", strconv.Itoa(cash.Supercategory), nil, errors.Get("supercategory")),
		Div(Class("d-flex justify-content-between"),
			c.Submit(cash.ID),
			If(cash.ID != 0,
				Button(Class("btn btn-outline-danger"),
					htmx.Delete(path.DeleteBackofficeCash(path.NewCurrency(cash.Currency), cash.ID)),
					htmx.Confirm("Are you sure?"),
					Text("Видалити"),
				),
			),
		),
	)
}
