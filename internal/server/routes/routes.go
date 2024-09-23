package routes

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
)

func New(tmpl *template.Template, assets embed.FS) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", handlers.NewIndexHandler(tmpl))

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /items", handlers.NewGetItemsHandler(tmpl))

	mux.Handle("GET /ping", handlers.NewPingHandler(tmpl))

	mux.Handle("GET /pong", handlers.NewPongHandler(tmpl))

	return mux
}
