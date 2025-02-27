package components

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *ItemsComponent) form(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	_ = categories

	return Form(
		Div(
			Label(Text("Date")),
			c.Errors("date", errors),
		),
		Div(
			Label(Text("Sum")),
			Input(Value(c.sum(item.Sum))),
			c.Errors("sum", errors),
		),
		Div(
			Label(Text("Категорія")),
			c.Errors("category", errors),
		),
		Div(
			Input(Value(item.Description)),
		),
		Button(
			Type("submit"),
			If(item.ID == 0, Text("Створити")),
			If(item.ID != 0, Text("Оновити")),
		),
	)
}
