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

func cashesPath() string {
	return "/backoffice/cashes/"
}

func cashPath(id int) string {
	return cashesPath() + strconv.Itoa(id)
}

func editCashPath(id int) string {
	return cashPath(id) + "/edit"
}
