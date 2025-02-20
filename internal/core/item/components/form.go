package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/pkg/validation"
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
			Label(
				Text("Sum"),
			),
			Input(
				Value(c.sum(item.Sum)),
				c.Errors("sum", errors),
			),
		),
		Div(
			Input(Value(item.Description)),
		),
	)
}
