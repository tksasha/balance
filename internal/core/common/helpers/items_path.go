package helpers

import (
	"net/http"
)

func ItemsPath(request *http.Request, year, month int) string {
	u := base(items, request)

	values := u.Query()

	var yearValue string
	if request != nil {
		yearValue = request.URL.Query().Get("year")
	}

	values.Set("year", Year(year, yearValue))

	var monthValue string
	if request != nil {
		monthValue = request.URL.Query().Get("month")
	}

	values.Set("month", Month(month, monthValue))

	u.RawQuery = values.Encode()

	return u.String()
}
