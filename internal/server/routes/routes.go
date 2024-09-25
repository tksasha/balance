package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/server/app"
)

func New(app *app.App, assets embed.FS) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", handlers.NewIndexHandler(app))

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /items", handlers.NewGetItemsHandler(app))

	mux.Handle("POST /items", handlers.NewCreateItemHandler(app))

	mux.Handle("GET /ping", handlers.NewPingHandler(app))

	mux.Handle("GET /pong", handlers.NewPongHandler(app))

	return mux
}
