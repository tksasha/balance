package component

import "time"

func date(date time.Time) string {
	if date.IsZero() {
		return ""
	}

	return date.Format(time.DateOnly)
}
