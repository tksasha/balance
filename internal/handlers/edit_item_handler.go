package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/errors"
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
		getItemService: services.NewGetItemService(app),
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

	if _, err := w.Write([]byte(item.Description)); err != nil {
		slog.Error(err.Error())
	}
}
