package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
)

type GetItemsHandler struct {
	itemService ItemService
}

func NewGetItemsHandler(itemService ItemService) *GetItemsHandler {
	return &GetItemsHandler{itemService}
}

func (h *GetItemsHandler) Pattern() string {
	return "GET /items"
}

func (h *GetItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("get items handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *GetItemsHandler) handle(w http.ResponseWriter, r *http.Request) error {
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
