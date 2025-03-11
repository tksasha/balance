package routes

import (
	"embed"
	"net/http"
	"time"

	balance "github.com/tksasha/balance/internal/app/balance/handler"
	cash "github.com/tksasha/balance/internal/app/cash/handlers"
	category "github.com/tksasha/balance/internal/app/category/handlers"
	index "github.com/tksasha/balance/internal/app/index/handler"
	item "github.com/tksasha/balance/internal/app/item/handlers"
	backofficecash "github.com/tksasha/balance/internal/backoffice/cash/handlers"
	backofficecategory "github.com/tksasha/balance/internal/backoffice/category/handlers"
	backofficeindex "github.com/tksasha/balance/internal/backoffice/index/handler"
)

//go:embed assets
var assets embed.FS

type Routes struct {
	Mux *http.ServeMux
}

func New( //nolint:funlen
	backofficeCashCreateHandler *backofficecash.CreateHandler,
	backofficeCashDeleteHandler *backofficecash.DeleteHandler,
	backofficeCashEditHandler *backofficecash.EditHandler,
	backofficeCashIndexHandler *backofficecash.IndexHandler,
	backofficeCashNewHandler *backofficecash.NewHandler,
	backofficeCashUpdateHandler *backofficecash.UpdateHandler,
	backofficeCategoryCreateHandler *backofficecategory.CreateHandler,
	backofficeCategoryDeleteHandler *backofficecategory.DeleteHandler,
	backofficeCategoryEditHandler *backofficecategory.EditHandler,
	backofficeCategoryListHandler *backofficecategory.ListHandler,
	backofficeCategoryUpdateHandler *backofficecategory.UpdateHandler,
	backofficeIndexHandler *backofficeindex.IndexHandler,
	balanceShowHandler *balance.ShowHandler,
	cashEditHandler *cash.EditHandler,
	cashIndexHandler *cash.IndexHandler,
	cashUpdateHandler *cash.UpdateHandler,
	categoryIndexHandler *category.IndexHandler,
	indexHandler *index.Handler,
	itemCreateHandler *item.CreateHandler,
	itemDeleteHandler *item.DeleteHandler,
	itemEditHandler *item.EditHandler,
	itemIndexHandler *item.IndexHandler,
	itemNewHandler *item.NewHandler,
	itemUpdateHandler *item.UpdateHandler,
) *Routes {
	mux := http.NewServeMux()

	mux.Handle("GET /assets/{$}", http.RedirectHandler("/", http.StatusMovedPermanently))

	mux.HandleFunc("GET /assets/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=31536000")

		w.Header().Set("Expires", time.Now().AddDate(1, 0, 0).Format(http.TimeFormat))

		http.FileServerFS(assets).ServeHTTP(w, r)
	})

	mux.Handle("GET /{$}", indexHandler)

	mux.Handle("GET /items/new", itemNewHandler)
	mux.Handle("POST /items", itemCreateHandler)
	mux.Handle("GET /items", itemIndexHandler)
	mux.Handle("GET /items/{id}/edit", itemEditHandler)
	mux.Handle("PATCH /items/{id}", itemUpdateHandler)
	mux.Handle("DELETE /items/{id}", itemDeleteHandler)

	mux.Handle("GET /cashes", cashIndexHandler)
	mux.Handle("GET /cashes/{id}/edit", cashEditHandler)
	mux.Handle("PATCH /cashes/{id}", cashUpdateHandler)

	mux.Handle("GET /categories", categoryIndexHandler)

	mux.Handle("GET /balance", balanceShowHandler)

	mux.Handle("GET /backoffice", backofficeIndexHandler)

	mux.Handle("GET /backoffice/cashes", backofficeCashIndexHandler)
	mux.Handle("GET /backoffice/cashes/new", backofficeCashNewHandler)
	mux.Handle("POST /backoffice/cash", backofficeCashCreateHandler)
	mux.Handle("DELETE /backoffice/cashes/{id}", backofficeCashDeleteHandler)
	mux.Handle("GET /backoffice/cashes/{id}/edit", backofficeCashEditHandler)
	mux.Handle("PATCH /backoffice/cashes/{id}", backofficeCashUpdateHandler)

	mux.Handle("GET /backoffice/categories", backofficeCategoryListHandler)
	mux.Handle("POST /backoffice/categories", backofficeCategoryCreateHandler)
	mux.Handle("DELETE /backoffice/categories/{id}", backofficeCategoryDeleteHandler)
	mux.Handle("GET /backoffice/categories/{id}/edit", backofficeCategoryEditHandler)
	mux.Handle("PATCH /backoffice/categories/{id}", backofficeCategoryUpdateHandler)

	return &Routes{
		Mux: mux,
	}
}
