package component

import "github.com/tksasha/balance/internal/common/component"

type Component struct {
	*component.Component
}

func New(component *component.Component) *Component {
	return &Component{
		Component: component,
	}
}
