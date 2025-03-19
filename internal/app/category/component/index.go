package component

import (
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Index(categories category.GroupedCategories, values url.Values) Node {
	nodes := []Node{}

	for _, key := range categories.Keys() {
		node := Div(Class("card consolidation"),
			Div(Class("card-body"),
				Div(Class("card-text"),
					Table(
						TBody(c.categories(categories[key], values)),
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

func (c *Component) categories(categories category.Categories, values url.Values) Node {
	var nodes []Node

	nodes = append(nodes, Map(categories, func(category *category.Category) Node {
		return c.category(category, values)
	}))

	if categories.HasMoreThanOne() {
		nodes = append(nodes, c.Summary(categories.Sum()))
	}

	return Group(nodes)
}

func (c *Component) category(category *category.Category, values url.Values) Node {
	params := path.Params{
		"category": strconv.Itoa(category.ID),
	}

	return Tr(
		Td(Class("name"),
			Text(category.Name),
		),
		Td(Class("sum"),
			Div(Class("link"),
				htmx.Get(path.Items(values, params)),
				htmx.Target("#items"),
				Text(c.Money(category.Sum)),
			),
		),
	)
}
