package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *IndexComponent) form( //nolint:funlen
	item *item.Item,
	categories category.Categories,
) Node {
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
				Value(
					c.Date(item.Date),
				),
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
				Value(
					c.sum(item.Sum),
				),
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
					Map(categories.Expense(), c.option),
				),
				OptGroup(
					Attr("Label", "Income"),
					Map(categories.Income(), c.option),
				),
			),
		),
	)
}

func (c *IndexComponent) option(category *category.Category) Node {
	return Option(
		Value(strconv.Itoa(category.ID)),
		Text(category.Name),
	)
}

func (c *IndexComponent) sum(sum float64) string {
	if sum == 0.0 {
		return ""
	}

	return c.Money(sum)
}
