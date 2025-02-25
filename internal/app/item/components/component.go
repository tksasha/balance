package components

import (
	"github.com/tksasha/balance/internal/common/component"
)

type ItemsComponent struct {
	*component.Component
}

func NewItemsComponent() *ItemsComponent {
	return &ItemsComponent{
		Component: component.New(),
	}
}

func (c *ItemsComponent) sum(sum float64) string {
	if sum == 0.0 {
		return ""
	}

	return c.Money(sum)
}
