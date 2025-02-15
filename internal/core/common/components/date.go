package components

import "time"

func Date(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(time.DateOnly)
}
