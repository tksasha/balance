package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
	"github.com/tksasha/balance/pkg/validation"
)

type UpdateHandler struct {
	itemService     item.Service
	categoryService category.Service
	itemsComponent  *components.ItemsComponent
}

func NewUpdateHandler(
	itemService item.Service,
	categoryService category.Service,
	itemsComponent *components.ItemsComponent,
) *UpdateHandler {
	return &UpdateHandler{
		itemService:     itemService,
		categoryService: categoryService,
		itemsComponent:  itemsComponent,
	}
}

func (h *UpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.handle(r)
	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	var verrors validation.Errors
	if errors.As(err, &verrors) {
		categories, err := h.categoryService.List(r.Context())
		if err != nil {
			handlers.SetError(w, err)

			return
		}

		err = h.itemsComponent.Update(item, categories, verrors).Render(w)

		handlers.SetError(w, err)
	}

	handlers.SetError(w, err)
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
