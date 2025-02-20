package components

import "github.com/tksasha/balance/internal/core/common/components"

type CashComponent struct {
	*components.BaseComponent
}

func NewCashComponent(baseComponent *components.BaseComponent) *CashComponent {
	return &CashComponent{
		BaseComponent: baseComponent,
	}
}
