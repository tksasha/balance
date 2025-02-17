package components

import "github.com/tksasha/balance/internal/core/common/components"

func sum(sum float64) string {
	if sum == 0.0 {
		return ""
	}

	return components.Money(sum)
}
