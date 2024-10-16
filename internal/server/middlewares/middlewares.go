package middlewares

import (
	"net/http"

	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/server/routes"
)

type Middlewares struct {
	app    *app.App
	routes *http.ServeMux
}

func New(app *app.App, routes *routes.Routes) *Middlewares {
	return &Middlewares{
		app:    app,
		routes: routes.GetServeMux(),
	}
}

func (m *Middlewares) GetHandler() http.Handler {
	return RecoveryMiddleware(
		NewCurrencyMiddleware(m.app).Wrap(m.routes),
	)
}
