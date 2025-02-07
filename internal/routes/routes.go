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
	cashCreateHandler *handlers.CashCreateHandler,
	cashEditHandler *handlers.CashEditHandler,
	categoryCreateHandler *handlers.CategoryCreateHandler,
	categoryDeleteHandler *handlers.CategoryDeleteHandler,
	categoryListHandler *handlers.CategoryListHandler,
	editCategoryHandler *handlers.EditCategoryHandler,
	getItemHandler *handlers.GetItemHandler,
	getItemsHandler *handlers.GetItemsHandler,
	indexPageHandler *handlers.IndexPageHandler,
	itemCreateHandler *handlers.ItemCreateHandler,
	updateCategoryHandler *handlers.UpdateCategoryHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", indexPageHandler)

	mux.Handle("POST /items", itemCreateHandler)
	mux.Handle("GET /items", getItemsHandler)
	mux.Handle("GET /items/{id}", getItemHandler)

	mux.Handle("POST /categories", categoryCreateHandler)
	mux.Handle("GET /categories", categoryListHandler)
	mux.Handle("GET /categories/{id}/edit", editCategoryHandler)
	mux.Handle("PATCH /categories/{id}", updateCategoryHandler)
	mux.Handle("DELETE /categories/{id}", categoryDeleteHandler)

	mux.Handle("POST /cash", cashCreateHandler)
	mux.Handle("GET /cashes/{id}/edit", cashEditHandler)

	return &Routes{
		Mux: mux,
	}
}
