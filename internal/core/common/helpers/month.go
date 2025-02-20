package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func Month(primaryValue int, value string) string {
	var err error

	month := primaryValue

	if month == 0 {
		month, err = strconv.Atoi(value)
		if err != nil || month == 0 {
			month = int(time.Now().Month())
		}
	}

	return fmt.Sprintf("%02d", month)
}
