package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/app/item/component"
	"github.com/tksasha/balance/internal/app/item/components"
	"github.com/tksasha/balance/internal/common"
	"github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/validation"
)

type UpdateHandler struct {
	*handler.Handler

	itemService     item.Service
	categoryService item.CategoryService
	itemsComponent  *components.ItemsComponent
	component       *component.Component
}

func NewUpdateHandler(
	itemService item.Service,
	categoryService item.CategoryService,
	itemsComponent *components.ItemsComponent,
) *UpdateHandler {
	return &UpdateHandler{
		Handler:         handler.New(),
		itemService:     itemService,
		categoryService: categoryService,
		itemsComponent:  itemsComponent,
		component:       component.New(),
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(r)
	if err == nil {
		err := h.component.Update(item).Render(w)

		h.SetError(w, err)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		categories, err := h.categoryService.List(r.Context())
		if err != nil {
			h.SetError(w, err)

			return
		}

		err = h.component.Edit(item, categories, verrors).Render(w)

		h.SetError(w, err)

		return
	}

	h.SetError(w, err)
}

func (h *UpdateHandler) handle(r *http.Request) (*item.Item, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := item.UpdateRequest{
		ID:          r.PathValue("id"),
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	return h.itemService.Update(r.Context(), request)
}
