package handlers

import (
	"log/slog"
	"net/http"

	indexcomponents "github.com/tksasha/balance/internal/components/index"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type IndexHandler struct {
	categoriesGetter services.CategoriesGetter
}

func NewIndexHandler(app *app.App) http.Handler {
	return &IndexHandler{
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoriesGetter.GetCategories(r.Context(), 0) // TODO: use currency instead
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := indexcomponents.IndexPage(categories).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
