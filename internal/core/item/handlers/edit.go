package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type EditHandler struct {
	*common.BaseHandler

	itemService     item.Service
	categoryService category.Service
	itemsComponent  *components.ItemsComponent
}

func NewEditHandler(
	baseHandler *common.BaseHandler,
	itemService item.Service,
	categoryService category.Service,
	itemsComponent *components.ItemsComponent,
) *EditHandler {
	return &EditHandler{
		BaseHandler:     baseHandler,
		itemService:     itemService,
		categoryService: categoryService,
		itemsComponent:  itemsComponent,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(w, r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	categories, err := h.categoryService.List(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.itemsComponent.Edit(item, categories).Render(w)

	h.SetError(w, err)
}

func (h *EditHandler) handle(_ http.ResponseWriter, r *http.Request) (*item.Item, error) {
	return h.itemService.Edit(r.Context(), r.PathValue("id"))
}
