package routes

import (
	"embed"
	"net/http"

	cash "github.com/tksasha/balance/internal/core/cash/handlers"
	category "github.com/tksasha/balance/internal/core/category/handlers"
	index "github.com/tksasha/balance/internal/core/index/handler"
	item "github.com/tksasha/balance/internal/core/item/handlers"
)

//go:embed assets
var assets embed.FS

type Routes struct {
	Mux *http.ServeMux
}

func New(
	cashCreateHandler *cash.CreateHandler,
	cashDeleteHandler *cash.DeleteHandler,
	cashEditHandler *cash.EditHandler,
	cashIndexHandler *cash.IndexHandler,
	cashNewHandler *cash.NewHandler,
	cashUpdateHandler *cash.UpdateHandler,
	categoryCreateHandler *category.CreateHandler,
	categoryDeleteHandler *category.DeleteHandler,
	categoryEditHandler *category.EditHandler,
	categoryIndexHandler *category.IndexHandler,
	categoryUpdateHandler *category.UpdateHandler,
	indexPageHandler *index.Handler,
	itemCreateHandler *item.CreateHandler,
	itemEditHandler *item.EditHandler,
	itemIndexHandler *item.IndexHandler,
	itemUpdateHandler *item.UpdateHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /...", indexPageHandler)

	mux.Handle("POST /items", itemCreateHandler)
	mux.Handle("GET /items", itemIndexHandler)
	mux.Handle("GET /items/{id}/edit", itemEditHandler)
	mux.Handle("PATCH /items/{id}", itemUpdateHandler)

	mux.Handle("POST /categories", categoryCreateHandler)
	mux.Handle("GET /categories", categoryIndexHandler)
	mux.Handle("GET /categories/{id}/edit", categoryEditHandler)
	mux.Handle("PATCH /categories/{id}", categoryUpdateHandler)
	mux.Handle("DELETE /categories/{id}", categoryDeleteHandler)

	mux.Handle("DELETE /cashes/{id}", cashDeleteHandler)
	mux.Handle("GET /cashes", cashIndexHandler)
	mux.Handle("GET /cashes/new", cashNewHandler)
	mux.Handle("GET /cashes/{id}/edit", cashEditHandler)
	mux.Handle("PATCH /cashes/{id}", cashUpdateHandler)
	mux.Handle("POST /cash", cashCreateHandler)

	return &Routes{
		Mux: mux,
	}
}
