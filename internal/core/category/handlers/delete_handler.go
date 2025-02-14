package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/responses"
)

type DeleteHandler struct {
	service category.Service
}

func NewDeleteHandler(service category.Service) *DeleteHandler {
	return &DeleteHandler{
		service: service,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.service.Delete(r.Context(), r.PathValue("id"))
}
