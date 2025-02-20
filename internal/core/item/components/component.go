package components

import "github.com/tksasha/balance/internal/core/common"

type ItemsComponent struct {
	*common.BaseComponent
}

func NewItemsComponent(baseComponent *common.BaseComponent) *ItemsComponent {
	return &ItemsComponent{
		BaseComponent: baseComponent,
	}
}
