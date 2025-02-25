package component

import (
	"time"

	"github.com/tksasha/balance/internal/common/component"
)

const yearFrom = 2015

type IndexComponent struct {
	*component.Component

	years []int
}

func NewIndexComponent() *IndexComponent {
	var years []int

	for year := yearFrom; year <= time.Now().Year(); year++ {
		years = append(years, year)
	}

	return &IndexComponent{
		Component: component.New(),
		years:     years,
	}
}
