package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Form(
	params params.Params,
	item *item.Item,
	categories category.Categories,
	errors validation.Errors,
) Node {
	return Form(
		If(item.ID == 0, htmx.Post(paths.CreateItem(params))),
		If(item.ID != 0, htmx.Patch(paths.UpdateItem(params, item.ID))),
		c.Input("Дата", "date", c.date(item.Date), nil, errors.Get("date")),
		c.Input("Сума", "formula", item.Formula, nil, errors.Get("sum"), AutoFocus()),
		If(
			item.CategoryVisible,
			Div(Class("mb-3"),
				c.categories(item.CategoryID, categories, errors.Get("category")),
			),
		),
		If(
			!item.CategoryVisible,
			Div(Class("mb-3"),
				Label(Class("form-label"), Text("Категорія")),
				Input(Type("hidden"), Name("category_id"), Value(strconv.Itoa(item.CategoryID))),
				Input(Class("form-control"), ReadOnly(), Value(item.CategoryName)),
			),
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
						htmx.Delete(paths.DeleteItem(params, item.ID)),
						htmx.Confirm("Are you sure?"),
						Text("Видалити"),
					),
				),
			),
		),
	)
}
