package components

import (
	"github.com/tksasha/balance/internal/core/common/component"
)

type ItemsComponent struct {
	*component.Component
}

func NewItemsComponent(component *component.Component) *ItemsComponent {
	return &ItemsComponent{
		Component: component,
	}
}
