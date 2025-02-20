package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func Year(primaryValue int, value string) string {
	var err error

	year := primaryValue

	if year == 0 {
		year, err = strconv.Atoi(value)
		if err != nil || year == 0 {
			year = time.Now().Year()
		}
	}

	return fmt.Sprintf("%04d", year)
}
