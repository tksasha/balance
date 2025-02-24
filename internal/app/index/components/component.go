package components

import (
	"github.com/tksasha/balance/internal/common/component"
)

type IndexComponent struct {
	*component.Component

	monthsComponent *MonthsComponent
	yearsComponent  *YearsComponent
}

func NewIndexComponent(
	component *component.Component,
	monthsComponent *MonthsComponent,
	yearsComponent *YearsComponent,
) *IndexComponent {
	return &IndexComponent{
		Component:       component,
		monthsComponent: monthsComponent,
		yearsComponent:  yearsComponent,
	}
}
