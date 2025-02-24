package components

import (
	"github.com/tksasha/balance/internal/common/component"
)

type ItemsComponent struct {
	*component.Component
}

func NewItemsComponent(component *component.Component) *ItemsComponent {
	return &ItemsComponent{
		Component: component,
	}
}

func (c *ItemsComponent) sum(sum float64) string {
	if sum == 0.0 {
		return ""
	}

	return c.Money(sum)
}
