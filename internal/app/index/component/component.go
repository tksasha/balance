package component

import (
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/common/component"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

const yearFrom = 2015

type Component struct {
	*component.Component

	years []int
}

func New() *Component {
	var years []int

	for year := yearFrom; year <= time.Now().Year(); year++ {
		years = append(years, year)
	}

	return &Component{
		Component: component.New(),
		years:     years,
	}
}

func (c *Component) id(name string, id int) Node {
	return ID(name + "-" + strconv.Itoa(id))
}

func (c *Component) editCashPath(id int) string {
	return "/cashes/" + strconv.Itoa(id) + "/edit"
}
