package valueobjects

import "strconv"

type Year struct {
	provider CurrentDateProvider
}

func NewYear(provider CurrentDateProvider) Year {
	return Year{
		provider: provider,
	}
}

func (y Year) Parse(primaryValue int, value string) int {
	if primaryValue != 0 {
		return primaryValue
	}

	year, err := strconv.Atoi(value)
	if err != nil || year == 0 {
		return y.provider.CurrentYear()
	}

	return year
}
