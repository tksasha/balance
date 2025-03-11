package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type EditHandler struct {
	*handler.Handler

	itemService     item.Service
	categoryService item.CategoryService
	component       *component.Component
}

func NewEditHandler(
	itemService item.Service,
	categoryService item.CategoryService,
) *EditHandler {
	return &EditHandler{
		Handler:         handler.New(),
		itemService:     itemService,
		categoryService: categoryService,
		component:       component.New(),
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.itemService.Edit(r.Context(), r.PathValue("id"))
	if err != nil {
		h.SetError(w, err)

		return
	}

	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	w.Header().Add("Hx-Trigger-After-Swap", "balance.item.edit")

	err = h.component.Edit(item, categories, nil).Render(w)

	h.SetError(w, err)
}
