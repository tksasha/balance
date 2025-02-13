package handlers

import (
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/category"
	"github.com/tksasha/balance/internal/category/components"
	"github.com/tksasha/balance/internal/common/handlers"
)

type ListHandler struct {
	service category.Service
}

func NewListHandler(service category.Service) *ListHandler {
	return &ListHandler{
		service: service,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.CategoryList(categories).Render(w); err != nil {
		slog.Error("render categories component error", "error", err)

		handlers.E(w, err)

		return
	}
}

func (h *ListHandler) handle(r *http.Request) (category.Categories, error) {
	return h.service.List(r.Context())
}
