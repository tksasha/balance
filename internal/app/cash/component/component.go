package component

import (
	"strconv"

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

func (c *Component) editPath(id int) string {
	return "/cashes/" + strconv.Itoa(id) + "/edit"
}
