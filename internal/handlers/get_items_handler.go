package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/services"
)

type getItemsHandler struct {
	itemService services.ItemService
}

func NewGetItemsHandler(itemService services.ItemService) Handler {
	return &getItemsHandler{itemService}
}

func (h *getItemsHandler) Pattern() string {
	return "GET /items"
}

func (h *getItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)

	if err := h.handle(w, r); err != nil {
		slog.Error("get items handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *getItemsHandler) handle(w http.ResponseWriter, r *http.Request) error {
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
