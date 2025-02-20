package components

import (
	"github.com/tksasha/balance/internal/core/common/components"
)

type IndexPageComponent struct {
	*components.BaseComponent

	monthsComponent *MonthsComponent
}

func NewIndexPageComponent(
	baseComponent *components.BaseComponent,
	monthsComponent *MonthsComponent,
) *IndexPageComponent {
	return &IndexPageComponent{
		BaseComponent:   baseComponent,
		monthsComponent: monthsComponent,
	}
}
