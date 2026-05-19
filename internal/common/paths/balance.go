package paths

import (
	"net/url"

	"github.com/tksasha/balance/internal/common/paths/params"
)

const (
	BalancePathKey = "balancePath"
)

func Balance(params params.Params) string {
	path := url.URL{
		Path:     "/balance",
		RawQuery: params.String(),
	}

	return path.String()
}
