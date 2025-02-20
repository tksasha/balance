package components

import (
	"github.com/tksasha/balance/internal/core/common/component"
)

type IndexPageComponent struct {
	*component.Component

	monthsComponent *MonthsComponent
}

func NewIndexPageComponent(component *component.Component, monthsComponent *MonthsComponent) *IndexPageComponent {
	return &IndexPageComponent{
		Component:       component,
		monthsComponent: monthsComponent,
	}
}
