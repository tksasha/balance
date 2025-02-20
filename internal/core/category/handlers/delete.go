package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type DeleteHandler struct {
	*handlers.BaseHandler

	categoryService category.Service
}

func NewDeleteHandler(
	baseHandler *handlers.BaseHandler,
	categoryService category.Service,
) *DeleteHandler {
	return &DeleteHandler{
		BaseHandler:     baseHandler,
		categoryService: categoryService,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		h.SetError(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.categoryService.Delete(r.Context(), r.PathValue("id"))
}
