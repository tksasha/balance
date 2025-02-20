package components

import "github.com/tksasha/balance/internal/core/common/helpers"

type BaseComponent struct {
	Helpers *helpers.Helpers
}

func NewBaseComponent(helpers *helpers.Helpers) *BaseComponent {
	return &BaseComponent{
		Helpers: helpers,
	}
}
