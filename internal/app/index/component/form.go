package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(item *item.Item, categories category.Categories) Node {
	return Div(Class("container mt-4"),
		Div(Class("row"),
			Div(Class("col"),
				Div(Class("card mb-3"),
					Div(Class("card-body"),
						Form(
							Div(Class("mb-3"),
								Label(Class("form-label"),
									Text("Date"),
								),
								Input(
									If(true, Class("form-control invalid")),
									If(true, Class("form-control")),
									Value(c.Date(item.Date)),
								),
							),
							Div(Class("mb-3"),
								Label(Class("form-label"),
									Text("Sum"),
								),
								Input(
									If(true, Class("form-control invalid")),
									If(true, Class("form-control")),
									Value(c.sum(item.Sum)),
								),
							),
							Div(Class("mb-3"),
								Label(Class("form-label"),
									Text("Category"),
								),
								Select(Class("form-select"),
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
						),
					),
				),
			),
		),
	)
}

func (c *Component) option(category *category.Category) Node {
	return Option(
		Value(strconv.Itoa(category.ID)),
		Text(category.Name),
	)
}

func (c *Component) sum(sum float64) string {
	if sum == 0.0 {
		return ""
	}

	return c.Money(sum)
}
