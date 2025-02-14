package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/index/components"
)

type Handler struct {
	categoryService category.Service
}

func NewHandler(categoryService category.Service) *Handler {
	return &Handler{
		categoryService: categoryService,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		handlers.E(w, err)

		return
	}
}

func (h *Handler) handle(w http.ResponseWriter, r *http.Request) error {
	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		return err
	}

	return components.Index(categories).Render(w)
}
