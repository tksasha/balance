package routes

import (
	"embed"
	"net/http"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/handlers/cash"
)

//go:embed assets
var assets embed.FS

type Routes struct {
	Mux *http.ServeMux
}

func New(
	cashCreateHandler *cash.CreateHandler,
	createCategoryHandler *handlers.CreateCategoryHandler,
	createItemHandler *handlers.CreateItemHandler,
	editCategoryHandler *handlers.EditCategoryHandler,
	getCategoriesHandler *handlers.GetCategoriesHandler,
	getItemHandler *handlers.GetItemHandler,
	getItemsHandler *handlers.GetItemsHandler,
	indexPageHandler *handlers.IndexPageHandler,
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

	mux.Handle("POST /cash", cashCreateHandler)

	return &Routes{
		Mux: mux,
	}
}
