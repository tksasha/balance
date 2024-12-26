package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
)

//go:embed assets
var assets embed.FS

func New(
	indexPageHandler *handlers.IndexPageHandler,
	createItemHandler *handlers.CreateItemHandler,
	getItemsHandler *handlers.GetItemsHandler,
	getItemHandler *handlers.GetItemHandler,
	getCategoriesHandler *handlers.GetCategoriesHandler,
) *http.ServeMux {
	routes := http.NewServeMux()

	routes.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	routes.Handle("GET /assets/", http.FileServerFS(assets))

	routes.Handle("GET /", indexPageHandler)

	routes.Handle("POST /items", createItemHandler)
	routes.Handle("GET /items", getItemsHandler)
	routes.Handle("GET /items/{id}", getItemHandler)

	routes.Handle("POST /categories", getCategoriesHandler)

	return routes
}
