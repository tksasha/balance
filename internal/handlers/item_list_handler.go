package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
)

type ItemListHandler struct {
	itemService ItemService
}

func NewItemListHandler(itemService ItemService) *ItemListHandler {
	return &ItemListHandler{itemService}
}

func (h *ItemListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("get items handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *ItemListHandler) handle(w http.ResponseWriter, r *http.Request) error {
	items, err := h.itemService.GetItems(r.Context())
	if err != nil {
		slog.Error("failed to get items", "error", err)

		return err
	}

	if err := components.Items(items).Render(w); err != nil {
		slog.Error("failed to render items", "error", err)

		return err
	}

	return nil
}
