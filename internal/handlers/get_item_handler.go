package handlers

import (
	"log/slog"
	"net/http"
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
		slog.Error("get item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *GetItemHandler) handle(_ http.ResponseWriter, r *http.Request) error {
	slog.Info("debug", "id", r.PathValue("id"))

	item, err := h.itemService.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		slog.Error("failed to get item by id", "error", err)

		return err
	}

	_ = item

	return nil
}
