package component

import (
	"github.com/tksasha/balance/internal/app/category"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Index(categories category.GroupedCategories) Node {
	nodes := []Node{}

	for _, key := range categories.Keys() {
		node := Div(Class("card consolidation"),
			Div(Class("card-body"),
				Div(Class("card-text"),
					Table(
						TBody(c.entities(categories[key])),
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

func (c *Component) entities(entities category.Categories) Node {
	var nodes []Node

	nodes = append(nodes, Map(entities, c.entity))

	if entities.HasMoreThanOne() {
		nodes = append(nodes, c.Summary(entities.Sum()))
	}

	return Group(nodes)
}

func (c *Component) entity(entity *category.Category) Node {
	return Tr(
		Td(Class("name"),
			Text(entity.Name),
		),
		Td(Class("sum"),
			Div(Class("link"),
				Text(c.Money(entity.Sum)),
			),
		),
	)
}
