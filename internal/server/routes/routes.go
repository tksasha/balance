package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/server/app"
)

func New(config *config.Config, app *app.App, assets embed.FS) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.Handle("GET /assets/", http.FileServerFS(assets))

	for _, currency := range config.Currencies {
		mux.Handle(
			"GET /"+currency.Name,
			handlers.NewIndexHandler(currency, app),
		)

		mux.Handle(
			"GET /"+currency.Name+"/items",
			handlers.NewGetItemsHandler(currency, app),
		)

		mux.Handle(
			"POST /"+currency.Name+"/items",
			handlers.NewCreateItemHandler(currency, app),
		)

		mux.Handle(
			"GET /"+currency.Name+"/items/{id}/edit",
			handlers.NewEditItemHandler(currency, app),
		)

		mux.Handle(
			"PATCH /"+currency.Name+"/items/{id}",
			handlers.NewUpdateItemHandler(currency, app),
		)
	}

	mux.Handle("DELETE /items/{id}", handlers.NewDeleteItemHandler(app))

	return mux
}
