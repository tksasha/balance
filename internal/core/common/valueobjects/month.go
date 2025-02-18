package valueobjects

import "strconv"

type Month struct {
	provider CurrentDateProvider
}

func NewMonth(provider CurrentDateProvider) Month {
	return Month{
		provider: provider,
	}
}

func (y Month) Parse(primaryValue int, value string) int {
	if primaryValue != 0 {
		return primaryValue
	}

	month, err := strconv.Atoi(value)
	if err != nil || month == 0 {
		return y.provider.CurrentMonth()
	}

	return month
}
