package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/requests"
)

type ItemUpdateHandler struct {
	itemService ItemService
}

func NewItemUpdateHandler(itemService ItemService) *ItemUpdateHandler {
	return &ItemUpdateHandler{
		itemService: itemService,
	}
}

func (h *ItemUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		if errors.Is(err, apperrors.ErrParsingForm) {
			http.Error(w, "Bad Request", http.StatusBadRequest)

			return
		}

		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)

			return
		}

		slog.Error("update item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("render update page\n"))
}

func (h *ItemUpdateHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return apperrors.ErrParsingForm
	}

	return h.itemService.Update(
		r.Context(),
		requests.UpdateItemRequest{
			ID:          r.PathValue("id"),
			Date:        r.FormValue("date"),
			Formula:     r.FormValue("formula"),
			CategoryID:  r.FormValue("category_id"),
			Description: r.FormValue("description"),
		},
	)
}
