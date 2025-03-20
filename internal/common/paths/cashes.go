package paths

import (
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/common/paths/params"
)

const cashesPath = "/cashes"

func Cashes(params params.Params) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: params.String(),
	}

	return path.String()
}

func EditCash(params params.Params, id int) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: params.String(),
	}

	return path.JoinPath(strconv.Itoa(id), "edit").String()
}

func UpdateCash(params params.Params, id int) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: params.String(),
	}

	return path.JoinPath(strconv.Itoa(id)).String()
}
