package component

import (
	"time"

	"github.com/tksasha/balance/internal/common/component"
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
