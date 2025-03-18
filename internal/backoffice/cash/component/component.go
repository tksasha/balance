package component

import (
	"github.com/tksasha/balance/internal/common/component"
)

type Component struct {
	*component.Component
}

func New() *Component {
	return &Component{
		Component: component.New(),
	}
}
