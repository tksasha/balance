package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	internalerrors "github.com/tksasha/balance/internal/errors"
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
		if errors.Is(err, internalerrors.ErrRecordNotFound) {
			http.Error(w, "Not Found", http.StatusNotFound)

			return
		}

		slog.Error("update item handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("render update page\n"))
}

func (h *UpdateItemHandler) handle(w http.ResponseWriter, r *http.Request) error {
	item, err := h.itemService.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		return err
	}

	_ = w

	return h.itemService.UpdateItem(r.Context(), item)
}
