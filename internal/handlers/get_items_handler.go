package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/services"
)

type getItemsHandler struct {
	itemService services.ItemService
}

func NewGetItemsHandler(itemService services.ItemService) Route {
	return &getItemsHandler{itemService}
}

func (h *getItemsHandler) Pattern() string {
	return "GET /items"
}

func (h *getItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("get items handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *getItemsHandler) handle(w http.ResponseWriter, r *http.Request) error {
	items, err := h.itemService.GetItems(r.Context())
	if err != nil {
		return err
	}

	if err := components.Items(items).Render(w); err != nil {
		return err
	}

	return nil
}
