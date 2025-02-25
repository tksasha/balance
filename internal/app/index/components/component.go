package components

import (
	"github.com/tksasha/balance/internal/common/component"
)

type IndexComponent struct {
	*component.Component

	monthsComponent *MonthsComponent
	yearsComponent  *YearsComponent
}

func NewIndexComponent(monthsComponent *MonthsComponent, yearsComponent *YearsComponent) *IndexComponent {
	return &IndexComponent{
		Component:       component.New(),
		monthsComponent: monthsComponent,
		yearsComponent:  yearsComponent,
	}
}
