package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) form(category *category.Category, _ validation.Errors) Node {
	return Form(
		If(category.ID == 0, htmx.Post(path.CreateBackofficeCategory(nil))),
		If(category.ID != 0, htmx.Patch(path.UpdateBackofficeCategory(nil, category.ID))),
		Div(Class("mb-3"),
			Label(Class("form-label"), Text("Валюта")),
			Select(Class("form-select"), Name("currency"),
				c.CurrencyOptions(currency.GetCode(category.Currency))),
		),
		c.Input("Назва", "name", category.Name, nil, nil, AutoFocus()),
		Div(Class("mb-3"),
			Div(Class("form-check"),
				Input(Class("form-check-input"), Type("checkbox"), ID("is-category-visible"),
					If(category.Visible, Checked()),
				),
				Label(Class("form-check-label"), For("is-category-visible"),
					Text("Відображається"),
				),
			),
		),
		c.Submit(category.ID),
	)
}
