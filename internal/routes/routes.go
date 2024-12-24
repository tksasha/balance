package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
)

//go:embed assets
var assets embed.FS

func New(handlers handlers.Handlers) *http.ServeMux {
	routes := http.NewServeMux()

	routes.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	routes.Handle("GET /assets/", http.FileServerFS(assets))

	for _, handler := range handlers {
		routes.Handle(handler.Pattern(), handler)
	}

	return routes
}
