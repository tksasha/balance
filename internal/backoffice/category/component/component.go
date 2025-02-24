package component

import "github.com/tksasha/balance/internal/common/component"

type CategoryComponent struct {
	*component.Component
}

func New(component *component.Component) *CategoryComponent {
	return &CategoryComponent{
		Component: component,
	}
}
