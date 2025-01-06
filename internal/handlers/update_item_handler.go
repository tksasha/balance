package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/models"
)

type UpdateItemHandler struct {
	itemService ItemService
}

func NewUpdateItemHandler(itemService ItemService) *UpdateItemHandler {
	return &UpdateItemHandler{
		itemService: itemService,
	}
}

func (h *UpdateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("update item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("render update page\n"))
}

func (h *UpdateItemHandler) handle(w http.ResponseWriter, r *http.Request) error {
	_ = w

	item := &models.Item{}

	return h.itemService.UpdateItem(r.Context(), item)
}
