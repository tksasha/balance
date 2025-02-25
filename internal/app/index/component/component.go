package component

import (
	"github.com/tksasha/balance/internal/common/component"
)

type IndexComponent struct {
	*component.Component
}

func NewIndexComponent() *IndexComponent {
	return &IndexComponent{
		Component: component.New(),
	}
}
