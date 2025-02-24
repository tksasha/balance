package routes

import (
	"embed"
	"net/http"

	cash "github.com/tksasha/balance/internal/app/cash/handlers"
	category "github.com/tksasha/balance/internal/app/category/handlers"
	index "github.com/tksasha/balance/internal/app/index/handler"
	item "github.com/tksasha/balance/internal/app/item/handlers"
	backofficeCategory "github.com/tksasha/balance/internal/backoffice/category/handlers"
)

//go:embed assets
var assets embed.FS

type Routes struct {
	Mux *http.ServeMux
}

func New(
	backofficeCategoryCreateHandler *backofficeCategory.CreateHandler,
	backofficeCategoryDeleteHandler *backofficeCategory.DeleteHandler,
	backofficeCategoryListHandler *backofficeCategory.ListHandler,
	cashCreateHandler *cash.CreateHandler,
	cashDeleteHandler *cash.DeleteHandler,
	cashEditHandler *cash.EditHandler,
	cashListHandler *cash.ListHandler,
	cashNewHandler *cash.NewHandler,
	cashUpdateHandler *cash.UpdateHandler,
	categoryEditHandler *category.EditHandler,
	categoryListHandler *category.ListHandler,
	categoryUpdateHandler *category.UpdateHandler,
	indexHandler *index.Handler,
	itemCreateHandler *item.CreateHandler,
	itemEditHandler *item.EditHandler,
	itemListHandler *item.ListHandler,
	itemUpdateHandler *item.UpdateHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))
	mux.Handle("GET /assets/", http.FileServerFS(assets))

	mux.Handle("GET /{$}", indexHandler)

	mux.Handle("POST /items", itemCreateHandler)
	mux.Handle("GET /items/{$}", itemListHandler)
	mux.Handle("GET /items/{id}/edit", itemEditHandler)
	mux.Handle("PATCH /items/{id}", itemUpdateHandler)

	mux.Handle("GET /categories", categoryListHandler)
	mux.Handle("GET /categories/{id}/edit", categoryEditHandler)
	mux.Handle("PATCH /categories/{id}", categoryUpdateHandler)

	mux.Handle("DELETE /cashes/{id}", cashDeleteHandler)
	mux.Handle("GET /cashes", cashListHandler)
	mux.Handle("GET /cashes/new", cashNewHandler)
	mux.Handle("GET /cashes/{id}/edit", cashEditHandler)
	mux.Handle("PATCH /cashes/{id}", cashUpdateHandler)
	mux.Handle("POST /cash", cashCreateHandler)

	mux.Handle("GET /backoffice/categories", backofficeCategoryListHandler)
	mux.Handle("POST /backoffice/categories", backofficeCategoryCreateHandler)
	mux.Handle("DELETE /backoffice/categories/{id}", backofficeCategoryDeleteHandler)

	return &Routes{
		Mux: mux,
	}
}
