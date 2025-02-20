package components

import "github.com/tksasha/balance/internal/core/common/components"

type CategoryComponent struct {
	*components.BaseComponent
}

func NewCategoryComponent(baseComponent *components.BaseComponent) *CategoryComponent {
	return &CategoryComponent{
		BaseComponent: baseComponent,
	}
}
