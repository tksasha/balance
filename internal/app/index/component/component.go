package component

import (
	balancecomponent "github.com/tksasha/balance/internal/app/balance/component"
	cashcomponent "github.com/tksasha/balance/internal/app/cash/component"
	"github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/balance/pkg/timeprovider"
)

const yearFrom = 2015

type Component struct {
	*component.Component

	balanceComponent *balancecomponent.Component
	cashComponent    *cashcomponent.Component
	years            []int
	timeProvider     timeprovider.TimeProvider
}

func New(timeProvider timeprovider.TimeProvider) *Component {
	var years []int

	currentYear := timeProvider.CurrentYear()

	for year := yearFrom; year <= currentYear; year++ {
		years = append(years, year)
	}

	return &Component{
		Component:        component.New(),
		balanceComponent: balancecomponent.New(),
		cashComponent:    cashcomponent.New(),
		years:            years,
		timeProvider:     timeProvider,
	}
}
