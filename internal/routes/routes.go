package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
)

//go:embed assets
var assets embed.FS

type Routes struct {
	Mux *http.ServeMux
}

func New(
	indexPageHandler *handlers.IndexPageHandler,
	createItemHandler *handlers.CreateItemHandler,
	getItemsHandler *handlers.GetItemsHandler,
	getItemHandler *handlers.GetItemHandler,
	createCategoryHandler *handlers.CreateCategoryHandler,
	getCategoriesHandler *handlers.GetCategoriesHandler,
	editCategoryHandler *handlers.EditCategoryHandler,
	updateCategoryHandler *handlers.UpdateCategoryHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", indexPageHandler)

	mux.Handle("POST /items", createItemHandler)
	mux.Handle("GET /items", getItemsHandler)
	mux.Handle("GET /items/{id}", getItemHandler)

	mux.Handle("POST /categories", createCategoryHandler)
	mux.Handle("GET /categories", getCategoriesHandler)
	mux.Handle("GET /categories/{id}/edit", editCategoryHandler)
	mux.Handle("PATCH /categories/{id}", updateCategoryHandler)

	return &Routes{
		Mux: mux,
	}
}
