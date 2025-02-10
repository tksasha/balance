package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
)

type ItemEditHandler struct {
	itemService ItemService
}

func NewItemEditHandler(itemService ItemService) *ItemEditHandler {
	return &ItemEditHandler{
		itemService: itemService,
	}
}

func (h *ItemEditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}

		slog.Error("get item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *ItemEditHandler) handle(_ http.ResponseWriter, r *http.Request) error {
	item, err := h.itemService.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		return err
	}

	_ = item

	return nil
}
