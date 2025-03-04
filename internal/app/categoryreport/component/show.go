package component

import (
	"github.com/tksasha/balance/internal/app/categoryreport"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Show(entities categoryreport.Entities) Node {
	return Map(entities, c.entity)
}

func (c *Component) entity(entity *categoryreport.Entity) Node {
	return Div(
		Div(Text(entity.CategoryName)),
		Div(Text(c.Money(entity.Sum))),
	)
}
