package components

import "github.com/tksasha/balance/internal/core/common"

type CashComponent struct {
	*common.BaseComponent
}

func NewCashComponent(baseComponent *common.BaseComponent) *CashComponent {
	return &CashComponent{
		BaseComponent: baseComponent,
	}
}
