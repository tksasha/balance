package components

import (
	"time"

	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/internal/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func form(item *item.Item, categories category.Categories) Node {
	return Form(
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Date"),
			),
			Input(
				If(
					true,
					Class("form-control invalid"),
				),
				If(
					true,
					Class("form-control"),
				),
				Value(item.Date.Format(time.DateOnly)),
			),
		),
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Sum"),
			),
			Input(
				If(
					true,
					Class("form-control invalid"),
				),
				If(
					true,
					Class("form-control"),
				),
				Value("sum should be here"),
			),
		),
		Div(
			Class("mb-3"),
			Label(
				Class("form-label"),
				Text("Category"),
			),
			Select(
				Class("form-select"),
				OptGroup(
					Attr("Label", "Expense"),
					Map(categories.Expense(), option),
				),
				OptGroup(
					Attr("Label", "Income"),
					Map(categories.Income(), option),
				),
			),
		),
	)
}
