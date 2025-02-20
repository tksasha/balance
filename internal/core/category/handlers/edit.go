package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type EditHandler struct {
	categoryService   category.Service
	categoryComponent *components.CategoryComponent
}

func NewEditHandler(
	categoryService category.Service,
	categoryComponent *components.CategoryComponent,
) *EditHandler {
	return &EditHandler{
		categoryService:   categoryService,
		categoryComponent: categoryComponent,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = h.categoryComponent.Edit(category).Render(w)

	handlers.SetError(w, err)
}

func (h *EditHandler) handle(r *http.Request) (*category.Category, error) {
	return h.categoryService.Edit(r.Context(), r.PathValue("id"))
}
