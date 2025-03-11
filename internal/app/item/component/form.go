package component

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *Component) Form(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	return Form(
		If(item.ID == 0, htmx.Post(path.CreateItem())),
		If(item.ID != 1, htmx.Patch(path.UpdateItem(item.ID))),
		c.Input("Дата", "date", c.date(item.Date), Classes{"datepicker": true}, errors.Get("date")),
		c.Input("Сума", "formula", item.Formula, nil, errors.Get("sum"), AutoFocus()),
		Div(Class("mb-3"),
			c.categories(item.CategoryID, categories, errors.Get("category")),
		),
		c.Input("Опис", "description", item.Description, nil, errors.Get("description")),
		Div(Class("form-group mt-4"),
			Div(Class("float-start"),
				Button(Class("btn btn-primary"),
					If(item.ID == 0, Text("Створити")),
					If(item.ID != 0, Text("Оновити")),
				),
			),
			If(item.ID != 0,
				Div(Class("float-end"),
					Button(Class("btn btn-outline-danger"),
						htmx.Delete(path.DeleteItem(item.ID)),
						htmx.Confirm("Are you sure?"),
						Text("Видалити"),
					),
				),
			),
		),
	)
}
