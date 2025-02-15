package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func form(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	_ = categories

	return Form(
		Div(
			Label(Text("Date")),
			components.Errors("date", errors),
		),
		Div(
			Label(
				Text("Sum"),
			),
			Input(
				Value(sum(item.Sum)),
				components.Errors("sum", errors),
			),
		),
		Div(
			Input(Value(item.Description)),
		),
	)
}
