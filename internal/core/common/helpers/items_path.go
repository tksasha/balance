package helpers

import (
	"fmt"
	"net/http"

	"github.com/tksasha/balance/internal/core/common/valueobjects"
)

func (h *Helpers) ItemsPath(request *http.Request, year, month int) string {
	u := base(items, request)

	values := u.Query()

	var yearValue string
	if request != nil {
		yearValue = request.URL.Query().Get("year")
	}

	parsedYear := valueobjects.NewYear(h.currentDateProvider).Parse(year, yearValue)

	values.Set("year", fmt.Sprintf("%04d", parsedYear))

	var monthValue string
	if request != nil {
		monthValue = request.URL.Query().Get("month")
	}

	parsedMonth := valueobjects.NewMonth(h.currentDateProvider).Parse(month, monthValue)

	values.Set("month", fmt.Sprintf("%02d", parsedMonth))

	u.RawQuery = values.Encode()

	return u.String()
}
