package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type EditHandler struct {
	*handler.Handler

	categoryService   category.Service
	categoryComponent *component.CategoryComponent
}

func NewEditHandler(
	categoryService category.Service,
	categoryComponent *component.CategoryComponent,
) *EditHandler {
	return &EditHandler{
		Handler:           handler.New(),
		categoryService:   categoryService,
		categoryComponent: categoryComponent,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.categoryComponent.Edit(category).Render(w)

	h.SetError(w, err)
}

func (h *EditHandler) handle(r *http.Request) (*category.Category, error) {
	return h.categoryService.Edit(r.Context(), r.PathValue("id"))
}
