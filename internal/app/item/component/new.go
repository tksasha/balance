package component

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) New(params params.Params, categories category.Categories, children ...Node) Node {
	return Form(Class("new_item"), htmx.Post(paths.CreateItem(params)),
		Div(Class("item-inline-form-date"),
			Input(Class("form-control"), Placeholder("Дата"), Name("date")),
		),
		Div(Class("item-inline-form-formula"),
			Input(Class("form-control"), Placeholder("Сума"), Name("formula")),
		),
		Div(Class("item-inline-form-category"),
			c.categories(0, categories, nil),
		),
		Div(Class("item-inline-form-description"),
			Input(Class("form-control"), Placeholder("Примітка"), Name("description")),
		),
		Div(Class("float-right"),
			Button(Class("btn btn-primary"), Text("Створити")),
		),
		Group(children),
	)
}
