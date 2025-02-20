package components

import "github.com/tksasha/balance/internal/core/common/components"

type ItemsComponent struct {
	components.BaseComponent
}

func NewItemsComponent() *ItemsComponent {
	return &ItemsComponent{}
}
