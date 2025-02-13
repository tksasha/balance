package routes

import (
	"embed"
	"net/http"

	cashhandlers "github.com/tksasha/balance/internal/cash/handlers"
	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/handlers/cashes"
)

//go:embed assets
var assets embed.FS

type Routes struct {
	Mux *http.ServeMux
}

func New(
	cashCreateHandler *cashhandlers.CreateHandler,
	cashDeleteHandler *handlers.CashDeleteHandler,
	cashEditHandler *handlers.CashEditHandler,
	cashListHandler *cashhandlers.ListHandler,
	cashNewHandler *cashes.NewHandler,
	cashUpdateHandler *handlers.CashUpdateHandler,
	categoryCreateHandler *handlers.CategoryCreateHandler,
	categoryDeleteHandler *handlers.CategoryDeleteHandler,
	categoryEditHandler *handlers.CategoryEditHandler,
	categoryListHandler *handlers.CategoryListHandler,
	categoryUpdateHandler *handlers.CategoryUpdateHandler,
	indexPageHandler *handlers.IndexPageHandler,
	itemCreateHandler *handlers.ItemCreateHandler,
	itemEditHandler *handlers.ItemEditHandler,
	itemListHandler *handlers.ItemListHandler,
	itemUpdateHandler *handlers.ItemUpdateHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /", indexPageHandler)

	mux.Handle("POST /items", itemCreateHandler)
	mux.Handle("GET /items", itemListHandler)
	mux.Handle("GET /items/{id}/edit", itemEditHandler)
	mux.Handle("PATCH /items/{id}", itemUpdateHandler)

	mux.Handle("POST /categories", categoryCreateHandler)
	mux.Handle("GET /categories", categoryListHandler)
	mux.Handle("GET /categories/{id}/edit", categoryEditHandler)
	mux.Handle("PATCH /categories/{id}", categoryUpdateHandler)
	mux.Handle("DELETE /categories/{id}", categoryDeleteHandler)

	mux.Handle("DELETE /cashes/{id}", cashDeleteHandler)
	mux.Handle("GET /cashes", cashListHandler)
	mux.Handle("GET /cashes/new", cashNewHandler)
	mux.Handle("GET /cashes/{id}/edit", cashEditHandler)
	mux.Handle("PATCH /cashes/{id}", cashUpdateHandler)
	mux.Handle("POST /cash", cashCreateHandler)

	return &Routes{
		Mux: mux,
	}
}
