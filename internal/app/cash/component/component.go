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

func (c *Component) cashUpdatePath(id int) string {
	return "/cashes/" + strconv.Itoa(id)
}

func (c *Component) cashEditPath(id int) string {
	return "/cashes/" + strconv.Itoa(id) + "/edit"
}

func (c *Component) cashID(id int) string {
	return "cash-" + strconv.Itoa(id)
}
