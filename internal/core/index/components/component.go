package components

import (
	"github.com/tksasha/balance/internal/core/common/component"
)

type IndexComponent struct {
	*component.Component

	monthsComponent *MonthsComponent
}

func NewIndexComponent(component *component.Component, monthsComponent *MonthsComponent) *IndexComponent {
	return &IndexComponent{
		Component:       component,
		monthsComponent: monthsComponent,
	}
}
