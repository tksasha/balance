package middlewares

import (
	"net/http"

	"github.com/tksasha/balance/internal/server/app"
)

func New(app *app.App, routes *http.ServeMux) http.Handler {
	return RecoveryMiddleware(
		NewCurrencyMiddleware(app).Wrap(routes),
	)
}
