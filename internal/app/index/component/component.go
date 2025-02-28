package component

import (
	"time"

	cashcomponent "github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/balance/internal/common/component"
)

const yearFrom = 2015

type Component struct {
	*component.Component

	cashComponent *cashcomponent.Component
	years         []int
}

func New() *Component {
	var years []int

	for year := yearFrom; year <= time.Now().Year(); year++ {
		years = append(years, year)
	}

	return &Component{
		Component:     component.New(),
		cashComponent: cashcomponent.New(),
		years:         years,
	}
}
