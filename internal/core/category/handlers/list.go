package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/common"
)

type ListHandler struct {
	*common.BaseHandler

	categoryService   category.Service
	categoryComponent *components.CategoryComponent
}

func NewListHandler(
	baseHandler *common.BaseHandler,
	categoryService category.Service,
	categoryComponent *components.CategoryComponent,
) *ListHandler {
	return &ListHandler{
		BaseHandler:       baseHandler,
		categoryService:   categoryService,
		categoryComponent: categoryComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.categoryComponent.List(categories).Render(w)

	h.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (category.Categories, error) {
	return h.categoryService.List(r.Context())
}
