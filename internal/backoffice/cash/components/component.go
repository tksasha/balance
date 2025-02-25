package components

import (
	"github.com/tksasha/balance/internal/common/component"
)

type CashComponent struct {
	*component.Component
}

func NewCashComponent() *CashComponent {
	return &CashComponent{
		Component: component.New(),
	}
}
