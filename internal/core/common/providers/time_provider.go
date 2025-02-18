package providers

import "time"

type TimeProvider struct{}

func NewTimeProvider() *TimeProvider {
	return &TimeProvider{}
}

func (p *TimeProvider) CurrentYear() int {
	return time.Now().Year()
}

func (p *TimeProvider) CurrentMonth() int {
	return int(time.Now().Month())
}
