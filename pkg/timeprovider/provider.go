package timeprovider

import "time"

func New() TimeProvider {
	return &provider{}
}

type provider struct{}

func (p *provider) IsCurrentMonth(month int) bool {
	return time.Now().Month() == time.Month(month)
}

func (p *provider) IsCurrentYear(year int) bool {
	return time.Now().Year() == year
}

func (p *provider) CurrentYear() int {
	return time.Now().Year()
}
