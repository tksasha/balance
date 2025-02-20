package components

import "github.com/tksasha/balance/internal/core/common"

type CategoryComponent struct {
	*common.BaseComponent
}

func NewCategoryComponent(baseComponent *common.BaseComponent) *CategoryComponent {
	return &CategoryComponent{
		BaseComponent: baseComponent,
	}
}
