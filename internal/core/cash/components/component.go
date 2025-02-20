package components

import (
	"github.com/tksasha/balance/internal/core/common/component"
)

type CashComponent struct {
	*component.Component
}

func NewCashComponent(component *component.Component) *CashComponent {
	return &CashComponent{
		Component: component,
	}
}
