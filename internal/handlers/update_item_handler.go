package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/requests"
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
	if err := h.handle(r); err != nil {
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

func (h *UpdateItemHandler) handle(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return internalerrors.ErrParsingForm
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
