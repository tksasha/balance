package component

import "github.com/tksasha/balance/internal/core/common/components"

type Component struct {
	*components.BaseComponent
}

func New(baseComponent *components.BaseComponent) *Component {
	return &Component{
		BaseComponent: baseComponent,
	}
}
