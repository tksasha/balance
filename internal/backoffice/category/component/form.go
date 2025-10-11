package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint: staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: staticcheck
)

func (c *Component) form(category *category.Category, errors validation.Errors) Node {
	return Form(
		If(category.ID == 0, htmx.Post(paths.CreateBackofficeCategory())),
		If(category.ID != 0, htmx.Patch(paths.UpdateBackofficeCategory(category.ID))),
		Div(Class("mb-3"),
			Label(Class("form-label"), Text("Валюта")),
			Select(Class("form-select"), Name("currency"),
				c.CurrencyOptions(currency.GetCode(category.Currency))),
		),
		c.Input("Назва", "name", category.Name, nil, errors.Get("name"), AutoFocus()),
		c.CheckBox("Надходження", "income", "is-category-income", category.Income),
		c.CheckBox("Відображається", "visible", "is-category-visible", category.Visible),
		c.Input("Група", "supercategory", strconv.Itoa(category.Supercategory), nil, errors.Get("supercategory")),
		c.Input("Номер в списку", "number", strconv.Itoa(category.Number), nil, errors.Get("number")),
		c.Submit(category.ID),
	)
}
