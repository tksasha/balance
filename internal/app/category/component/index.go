package component

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Index(categories category.GroupedCategories, params params.Params) Node {
	nodes := []Node{}

	for _, key := range categories.Keys() {
		node := Div(Class("card consolidation"),
			Div(Class("card-body"),
				Div(Class("card-text"),
					Table(
						TBody(c.categories(categories[key], params)),
					),
				),
			),
		)

		nodes = append(nodes, node)
	}

	return Div(Class("container-fluid"), ID("categories"),
		Div(Class("clearfix mt-4"),
			Group(nodes),
		),
	)
}

func (c *Component) categories(categories category.Categories, params params.Params) Node {
	var nodes []Node

	nodes = append(nodes, Map(categories, func(category *category.Category) Node {
		return c.category(category, params)
	}))

	if categories.HasMoreThanOne() {
		nodes = append(nodes, c.Summary(categories.Sum()))
	}

	return Group(nodes)
}

func (c *Component) category(category *category.Category, params params.Params) Node {
	params = params.WithCategory(category.ID)

	return Tr(
		Td(Class("name"),
			Text(category.Name),
		),
		Td(Class("sum"),
			Div(Class("link"),
				htmx.Get(paths.Items(params)),
				htmx.Target("#items"),
				Text(c.Money(category.Sum)),
			),
		),
	)
}
