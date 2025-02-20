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

type CreateHandler struct {
	service         item.Service
	categoryService category.Service
	itemsComponent  *components.ItemsComponent
}

func NewCreateHandler(
	service item.Service,
	categoryService category.Service,
	itemsComponent *components.ItemsComponent,
) *CreateHandler {
	return &CreateHandler{
		service:         service,
		categoryService: categoryService,
		itemsComponent:  itemsComponent,
	}
}

func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

		err = h.itemsComponent.Create(item, categories, verrors).Render(w)

		handlers.SetError(w, err)

		return
	}

	handlers.SetError(w, err)
}

func (h *CreateHandler) handle(r *http.Request) (*item.Item, error) {
	if err := r.ParseForm(); err != nil {
		return nil, common.ErrParsingForm
	}

	request := item.CreateRequest{
		Date:        r.FormValue("date"),
		Formula:     r.FormValue("formula"),
		CategoryID:  r.FormValue("category_id"),
		Description: r.FormValue("description"),
	}

	return h.service.Create(r.Context(), request)
}
