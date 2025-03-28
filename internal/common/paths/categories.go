package paths

import (
	"net/url"
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/common/paths/params"
)

const categoriesPath = "/categories"

func Categories(params params.Params) string {
	if !params.Has("month") {
		month := int(time.Now().Month())

		params = params.With("month", strconv.Itoa(month))
	}

	if !params.Has("year") {
		year := time.Now().Year()

		params = params.With("year", strconv.Itoa(year))
	}

	path := url.URL{
		Path:     categoriesPath,
		RawQuery: params.String(),
	}

	return path.String()
}
