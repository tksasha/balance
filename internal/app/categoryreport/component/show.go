package component

import (
	"github.com/tksasha/balance/internal/app/categoryreport"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Show(entities categoryreport.MappedEntities) Node {
	nodes := []Node{}

	for _, key := range entities.Keys() {
		node := Div(Class("card consolidation"),
			Div(Class("card-body"),
				Div(Class("card-text"),
					Table(
						TBody(c.entities(entities[key])),
					),
				),
			),
		)

		nodes = append(nodes, node)
	}

	return Div(Class("container-fluid"),
		Div(Class("clearfix mt-4"),
			c.Map(nodes),
		),
	)
}

func (c *Component) entities(entities categoryreport.Entities) Node {
	var nodes []Node

	nodes = append(nodes, Map(entities, c.entity))

	if entities.HasMoreThanOne() {
		nodes = append(nodes, c.Summary(entities.Sum()))
	}

	return c.Map(nodes)
}

func (c *Component) entity(entity *categoryreport.Entity) Node {
	return Tr(
		Td(Class("name"),
			Text(entity.CategoryName),
		),
		Td(Class("sum"),
			Div(Class("link"),
				Text(c.Money(entity.Sum)),
			),
		),
	)
}
