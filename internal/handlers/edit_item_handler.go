package handlers

import (
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	"github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type EditItemHandler struct {
	getItemService *services.GetItemService
}

func NewEditItemHandler(
	app *app.App,
) http.Handler {
	return &EditItemHandler{
		getItemService: services.NewGetItemService(
			repositories.NewItemRepository(app.DB),
		),
	}
}

func (h *EditItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.getItemService.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		var notFoundError *errors.NotFoundError

		if errors.As(err, &notFoundError) {
			slog.Error(err.Error())

			w.WriteHeader(http.StatusNotFound)

			return
		}

		var unknownError *errors.UnknownError

		if errors.As(err, unknownError) {
			slog.Error(err.Error())

			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	}

	if err := itemcomponents.EditPage(item).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
