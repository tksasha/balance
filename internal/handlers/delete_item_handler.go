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

type DeleteItemHandler struct {
	itemGetter  services.ItemGetter
	itemDeleter services.ItemDeleter
}

func NewDeleteItemHandler(app *app.App) http.Handler {
	itemRepository := repositories.NewItemRepository(app.DB)

	return &DeleteItemHandler{
		itemGetter:  services.NewGetItemService(itemRepository),
		itemDeleter: services.NewDeleteItemService(itemRepository),
	}
}

func (h *DeleteItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.itemGetter.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		if errors.Is(err, internalerrors.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := h.itemDeleter.DeleteItem(r.Context(), item); err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := itemcomponents.DeletePage(item).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
