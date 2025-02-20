package components

import "github.com/tksasha/balance/internal/core/common"

type IndexPageComponent struct {
	*common.BaseComponent

	monthsComponent *MonthsComponent
}

func NewIndexPageComponent(
	baseComponent *common.BaseComponent,
	monthsComponent *MonthsComponent,
) *IndexPageComponent {
	return &IndexPageComponent{
		BaseComponent:   baseComponent,
		monthsComponent: monthsComponent,
	}
}
