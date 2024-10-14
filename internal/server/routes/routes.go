package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/config"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/server/app"
)

//go:embed assets
var assets embed.FS

func New(config *config.Config, app *app.App) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", handlers.NewIndexHandler(app))

	mux.Handle("GET /items", handlers.NewGetItemsHandler(app))

	mux.Handle("POST /items", handlers.NewCreateItemHandler(app))

	mux.Handle("GET /items/{id}/edit", handlers.NewEditItemHandler(app))

	mux.Handle("PATCH /items/{id}", handlers.NewUpdateItemHandler(app))

	mux.Handle("DELETE /items/{id}", handlers.NewDeleteItemHandler(app))

	return mux
}
