package component

import (
	"strconv"

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

func (c *Component) editPath(id int) string {
	return "/items" + strconv.Itoa(id) + "/edit"
}
