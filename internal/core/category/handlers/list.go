package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type ListHandler struct {
	categoryService   category.Service
	categoryComponent *components.CategoryComponent
}

func NewListHandler(
	categoryService category.Service,
	categoryComponent *components.CategoryComponent,
) *ListHandler {
	return &ListHandler{
		categoryService:   categoryService,
		categoryComponent: categoryComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = h.categoryComponent.List(categories).Render(w)

	handlers.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (category.Categories, error) {
	return h.categoryService.List(r.Context())
}
