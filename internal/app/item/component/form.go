package component

import (
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	return Form(
		c.Input("Дата", "date", item.Date.Format(time.DateOnly), errors.Get("date")),
		c.Input("Сума", "formula", item.Formula, errors.Get("sum")),
		Div(Class("mb-3"),
			c.categories(item.CategoryID, categories, errors.Get("category")),
		),
		c.Input("Опис", "description", item.Description, errors.Get("description")),
		c.Submit(item.ID),
	)
}

func (c *Component) categories(selected int, categories category.Categories, message *string) Node {
	var nodes []Node

	nodes = append(
		nodes,
		Select(Class("form-select"), Name("category"),
			OptGroup(Label(Text("Видатки")),
				Map(categories.Expense(), func(category *category.Category) Node {
					return c.category(category, selected)
				}),
			),
			OptGroup(Label(Text("Надходження")),
				Map(categories.Income(), func(category *category.Category) Node {
					return c.category(category, selected)
				}),
			),
		),
	)
	if message != nil {
		nodes = append(nodes, Div(Class("invalid-feebback"), Text(*message)))
	}

	return c.Map(nodes)
}

func (c *Component) category(category *category.Category, selected int) Node {
	return Option(
		Value(strconv.Itoa(category.ID)),
		Text(category.Name),
		If(category.ID == selected, Selected()),
	)
}
