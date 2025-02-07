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
	categoryEditHandler *handlers.CategoryEditHandler,
	categoryListHandler *handlers.CategoryListHandler,
	categoryUpdateHandler *handlers.CategoryUpdateHandler,
	indexPageHandler *handlers.IndexPageHandler,
	itemCreateHandler *handlers.ItemCreateHandler,
	itemEditHandler *handlers.ItemEditHandler,
	itemListHandler *handlers.ItemListHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", indexPageHandler)

	mux.Handle("POST /items", itemCreateHandler)
	mux.Handle("GET /items", itemListHandler)
	mux.Handle("GET /items/{id}/edit", itemEditHandler)

	mux.Handle("POST /categories", categoryCreateHandler)
	mux.Handle("GET /categories", categoryListHandler)
	mux.Handle("GET /categories/{id}/edit", categoryEditHandler)
	mux.Handle("PATCH /categories/{id}", categoryUpdateHandler)
	mux.Handle("DELETE /categories/{id}", categoryDeleteHandler)

	mux.Handle("POST /cash", cashCreateHandler)
	mux.Handle("GET /cashes/{id}/edit", cashEditHandler)

	return &Routes{
		Mux: mux,
	}
}
