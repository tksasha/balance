package component

import (
	"time"

	balancecomponent "github.com/tksasha/balance/internal/app/balance/component"
	cashcomponent "github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/balance/internal/common/component"
)

const yearFrom = 2015

type Component struct {
	*component.Component

	balanceComponent *balancecomponent.Component
	cashComponent    *cashcomponent.Component
	years            []int
}

func New() *Component {
	var years []int

	for year := yearFrom; year <= time.Now().Year(); year++ {
		years = append(years, year)
	}

	return &Component{
		Component:        component.New(),
		balanceComponent: balancecomponent.New(),
		cashComponent:    cashcomponent.New(),
		years:            years,
	}
}
