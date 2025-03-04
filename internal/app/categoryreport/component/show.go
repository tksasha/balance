package component

import (
	"maps"
	"slices"

	"github.com/tksasha/balance/internal/app/categoryreport"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Show(entities categoryreport.MappedEntities) Node {
	keys := slices.Collect(maps.Keys(entities))

	slices.Sort(keys)

	nodes := []Node{}

	for _, key := range keys {
		node := Div(Class("card consolidation"),
			Div(Class("card-body"),
				Div(Class("card-text"),
					Table(
						TBody(Map(entities[key], c.entity)),
					),
				),
			),
		)

		nodes = append(nodes, node)
	}

	return Div(Class("container-fluid"),
		Div(Class("clearfix mt-4"),
			Map(nodes, func(node Node) Node { return node }),
		),
	)
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
