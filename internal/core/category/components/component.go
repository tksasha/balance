package components

import (
	"github.com/tksasha/balance/internal/common/component"
)

type CategoryComponent struct {
	*component.Component
}

func NewCategoryComponent(component *component.Component) *CategoryComponent {
	return &CategoryComponent{
		Component: component,
	}
}
