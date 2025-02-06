package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	internalerrors "github.com/tksasha/balance/internal/errors"
)

type GetItemHandler struct {
	itemService ItemService
}

func NewGetItemHandler(itemService ItemService) *GetItemHandler {
	return &GetItemHandler{
		itemService: itemService,
	}
}

func (h *GetItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		if errors.Is(err, internalerrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}

		slog.Error("get item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *GetItemHandler) handle(_ http.ResponseWriter, r *http.Request) error {
	item, err := h.itemService.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		slog.Error("failed to get item by id", "error", err)

		return err
	}

	_ = item

	return nil
}
