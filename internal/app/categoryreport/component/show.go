package component

import (
	"github.com/tksasha/balance/internal/app/categoryreport"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Show(entities categoryreport.Entities) Node {
	return Div(Class("container-fluid"),
		Div(Class("clearfix mt-4"),
			Div(Class("card consolidation"),
				Div(Class("card-body"),
					Div(Class("card-text"),
						Table(
							TBody(
								Map(entities, c.entity),
							),
						),
					),
				),
			),
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
