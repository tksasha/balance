package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
)

type DeleteHandler struct {
	itemService item.Service
}

func NewDeleteHandler(itemService item.Service) *DeleteHandler {
	return &DeleteHandler{
		itemService: itemService,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		handlers.SetError(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.itemService.Delete(r.Context(), r.PathValue("id"))
}
