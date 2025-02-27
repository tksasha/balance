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
