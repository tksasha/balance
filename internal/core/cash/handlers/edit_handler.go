package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/responses"
)

type EditHandler struct {
	service cash.Service
}

func NewEditHandler(service cash.Service) *EditHandler {
	return &EditHandler{
		service: service,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}

	_, _ = w.Write([]byte(cash.Name))
}

func (h *EditHandler) handle(r *http.Request) (*cash.Cash, error) {
	return h.service.FindByID(r.Context(), r.PathValue("id"))
}
