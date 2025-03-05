package component

import (
	commoncomponent "github.com/tksasha/balance/internal/common/component"
)

type Component struct {
	*commoncomponent.Component
}

func New() *Component {
	return &Component{
		Component: commoncomponent.New(),
	}
}
