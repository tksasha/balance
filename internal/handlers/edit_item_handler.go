package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	internalerrors "github.com/tksasha/balance/internal/errors"
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
		if errors.Is(err, internalerrors.ErrNotFound) {
			slog.Error(err.Error())

			w.WriteHeader(http.StatusNotFound)

			return
		}

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := itemcomponents.EditPage(item).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
